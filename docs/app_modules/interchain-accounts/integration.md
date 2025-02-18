<!--
order: 2
-->

# Integrating Interchain Accounts into a Chain

Learn how to integrate Interchain Accounts host and controller functionality to your chain. The following document only applies for Cosmos SDK chains.

The Interchain Accounts module contains two submodules. Each submodule has its own IBC application. The Interchain Accounts module should be registered as an `AppModule` in the same way all SDK modules are registered on a chain, but each submodule should create its own `IBCModule` as necessary. A route should be added to the IBC router for each submodule which will be used. 

Chains who wish to support ICS27 may elect to act as a host chain, a controller chain or both. Disabling host or controller functionality may be done statically by excluding the host or controller module entirely from the `app.go` file or it may be done dynamically by taking advantage of the on-chain parameters which enable or disable the host or controller submodules. 

Interchain Account authentication modules are the base application of a middleware stack. The controller submodule is the middleware in this stack.


### Example integration

```go
// app.go

// Register the AppModule for the Interchain Accounts module and the authentication module
// Note: No `icaauth` exists, this must be substituted with an actual Interchain Accounts authentication module
ModuleBasics = module.NewBasicManager(
    ...
    ica.AppModuleBasic{},
    icaauth.AppModuleBasic{},
    ...
)

... 

// Add module account permissions for the Interchain Accounts module
// Only necessary for host chain functionality
// Each Interchain Account created on the host chain is derived from the module account created
maccPerms = map[string][]string{
    ...
    icatypes.ModuleName:            nil,
}

...

// Add Interchain Accounts Keepers for each submodule used and the authentication module
// If a submodule is being statically disabled, the associated Keeper does not need to be added. 
type App struct {
    ...

    ICAControllerKeeper icacontrollerkeeper.Keeper
    ICAHostKeeper       icahostkeeper.Keeper
    ICAAuthKeeper       icaauthkeeper.Keeper

    ...
}

...

// Create store keys for each submodule Keeper and the authentication module
keys := sdk.NewKVStoreKeys(
    ...
    icacontrollertypes.StoreKey,
    icahosttypes.StoreKey,
    icaauthtypes.StoreKey,
    ...
)

... 

// Create the scoped keepers for each submodule keeper and authentication keeper
scopedICAControllerKeeper := app.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
scopedICAHostKeeper := app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
scopedICAAuthKeeper := app.CapabilityKeeper.ScopeToModule(icaauthtypes.ModuleName)

...

// Create the Keeper for each submodule
app.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
		appCodec, keys[icacontrollertypes.StoreKey], app.GetSubspace(icacontrollertypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper, // may be replaced with middleware such as ics29 fee
		app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		app.AccountKeeper, scopedICAControllerKeeper, app.MsgServiceRouter(),
)
app.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec, keys[icahosttypes.StoreKey], app.GetSubspace(icahosttypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		app.AccountKeeper, scopedICAHostKeeper, app.MsgServiceRouter(),
)

// Create Interchain Accounts AppModule
icaModule := ica.NewAppModule(&app.ICAControllerKeeper, &app.ICAHostKeeper)

// Create your Interchain Accounts authentication module
app.ICAAuthKeeper = icaauthkeeper.NewKeeper(appCodec, keys[icaauthtypes.StoreKey], app.ICAControllerKeeper, scopedICAAuthKeeper)

// ICA auth AppModule
icaAuthModule := icaauth.NewAppModule(appCodec, app.ICAAuthKeeper)

// ICA auth IBC Module
ICAAuthIBCModule := icaauth.NewIBCModule(app.ICAAuthKeeper)

// Create host and controller IBC Modules as desired
icaControllerIBCModule := icacontroller.NewIBCModule(app.ICAControllerKeeper, icaAuthIBCModule)
icaHostIBCModule := icahost.NewIBCModule(app.ICAHostKeeper)

// Register host and authentication routes
ibcRouter.AddRoute(icacontrollertypes.SubModuleName, icaControllerIBCModule).
		AddRoute(icahosttypes.SubModuleName, icaHostIBCModule).
		AddRoute(icaauthtypes.ModuleName, icaControllerIBCModule) // Note, the authentication module is routed to the top level of the middleware stack

...

// Register Interchain Accounts and authentication module AppModule's
app.moduleManager = module.NewManager(
    ...
    icaModule,
    icaAuthModule,
)

...

// Add Interchain Accounts module InitGenesis logic
app.mm.SetOrderInitGenesis(
    ...
    icatypes.ModuleName,
    ...
)
```

## Parameters

The Interchain Accounts module contains the following on-chain parameters, logically separated for each distinct submodule:

### Controller Submodule Parameters

| Key                    | Type | Default Value |
|------------------------|------|---------------|
| `ControllerEnabled`    | bool | `true`        |

#### ControllerEnabled

The `ControllerEnabled` parameter controls a chains ability to service ICS-27 controller specific logic. This includes the sending of Interchain Accounts packet data as well as the following ICS-26 callback handlers:
- `OnChanOpenInit`
- `OnChanOpenAck`
- `OnChanCloseConfirm`
- `OnAcknowledgementPacket`
- `OnTimeoutPacket`

### Host Submodule Parameters

| Key                    | Type     | Default Value |
|------------------------|----------|---------------|
| `HostEnabled`          | bool     | `true`        |
| `AllowMessages`        | []string | `[]`          |

#### HostEnabled

The `HostEnabled` parameter controls a chains ability to service ICS27 host specific logic. This includes the following ICS-26 callback handlers:
- `OnChanOpenTry`
- `OnChanOpenConfirm`
- `OnChanCloseConfirm`
- `OnRecvPacket`

#### AllowMessages

The `AllowMessages` parameter provides the ability for a chain to limit the types of messages or transactions that it chooses to facilitate by defining an allowlist using the Protobuf message TypeURL format.

For example, a Cosmos SDK based chain that elects to provide hosted Interchain Accounts with the ability of governance voting and staking delegations will define its parameters as follows:

```
"params": {
    "host_enabled": true,
    "allow_messages": ["/cosmos.staking.v1beta1.MsgDelegate", "/cosmos.gov.v1beta1.MsgVote"]
}
```