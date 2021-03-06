package volcli

import "github.com/codegangsta/cli"

// GlobalFlags are required global flags for the operation of volcli.
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "prefix",
		Usage: "prefix key used in etcd for namespacing",
		Value: "/volplugin",
	},
	cli.StringSliceFlag{
		Name:  "etcd",
		Usage: "URL for etcd",
		Value: &cli.StringSlice{"http://localhost:2379"},
	},
	cli.StringFlag{
		Name:  "apiserver",
		Usage: "address of apiserver process",
		Value: "127.0.0.1:9005",
	},
}

// Commands is the data structure which describes the command hierarchy
// for volcli.
var Commands = []cli.Command{
	{
		Name:  "global",
		Usage: "Manage Global Configuration",
		Subcommands: []cli.Command{
			{
				Name:        "upload",
				ArgsUsage:   "accepts configuration from stdin",
				Description: "Uploads a the global configuration to etcd. Accepts JSON.",
				Usage:       "Upload a policy to etcd",
				Action:      GlobalUpload,
			},
			{
				Name:        "get",
				ArgsUsage:   "[policy name]",
				Usage:       "Obtain the global configuration",
				Description: "Gets the global configuration from etcd",
				Action:      GlobalGet,
			},
		},
	},
	{
		Name:  "policy",
		Usage: "Manage Policies",
		Subcommands: []cli.Command{
			{
				Name:        "upload",
				ArgsUsage:   "[policy name]. accepts from stdin",
				Description: "Uploads a policy to etcd. Accepts JSON. Requires direct, unauthenticated access to etcd.",
				Usage:       "Upload a policy to etcd",
				Action:      PolicyUpload,
			},
			{
				Name:        "delete",
				ArgsUsage:   "[policy name]",
				Description: "Permanently removes a policy from etcd. Volumes that belong to the policy are unaffected.",
				Usage:       "Delete a policy",
				Action:      PolicyDelete,
			},
			{
				Name:        "get",
				ArgsUsage:   "[policy name]",
				Usage:       "Obtain the policy",
				Description: "Gets the policy for a policy from etcd.",
				Action:      PolicyGet,
			},
			{
				Name:        "list",
				ArgsUsage:   "",
				Description: "Reads policies and generates a newline-delimited list",
				Usage:       "List all policies",
				Action:      PolicyList,
			},
			{
				Name:        "history",
				Description: "See historical policy revisions",
				Usage:       "See historical policy revisions",
				Subcommands: []cli.Command{
					{
						Name:        "get",
						ArgsUsage:   "[policy name] [revision]",
						Description: "Retrieve a single revision of a policy",
						Usage:       "Retrieve a single revision of a policy",
						Action:      PolicyGetRevision,
					},
					{
						Name:        "list",
						ArgsUsage:   "[policy name]",
						Description: "List all revisions of a policy",
						Usage:       "List all revisions of a policy",
						Action:      PolicyListRevisions,
					},
				},
			},
			{
				Name:        "watch",
				ArgsUsage:   "",
				Description: "Watch for policy changes",
				Usage:       "Watch for policy changes",
				Action:      PolicyWatch,
			},
		},
	},
	{
		Name:  "volume",
		Usage: "Manage Volumes",
		Subcommands: []cli.Command{
			{
				Name: "create",
				Flags: []cli.Flag{cli.StringSliceFlag{
					Name:  "opt",
					Usage: "Provide key=value options to create the volume",
				}},
				ArgsUsage:   "[policy name]/[volume name]",
				Description: "This creates a logical volume. Calls out to the apiserver and sets the policy based on the policy name provided.",
				Usage:       "Create a volume for a given policy",
				Action:      VolumeCreate,
			},
			{
				Name:        "get",
				ArgsUsage:   "[policy name]/[volume name]",
				Usage:       "Get JSON configuration for a volume",
				Description: "Obtain the JSON configuration for the volume",
				Action:      VolumeGet,
			},
			{
				Name:        "list",
				ArgsUsage:   "[policy name]",
				Description: "Given a policy name, produces a newline-delimited list of volumes.",
				Usage:       "List all volumes for a given policy",
				Action:      VolumeList,
			},
			{
				Name:        "list-all",
				ArgsUsage:   "",
				Description: "Produces a newline-delimited list of policy/volume combinations.",
				Usage:       "List all volumes across policies",
				Action:      VolumeListAll,
			},
			{
				Name:        "force-remove",
				ArgsUsage:   "[policy name]/[volume name]",
				Description: "Forcefully removes a volume without deleting or unmounting the underlying image",
				Usage:       "Forcefully remove a volume without removing the underlying image",
				Action:      VolumeForceRemove,
			},
			{
				Name: "remove",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "timeout, t",
						Usage: "Retry to remove the volume until the timeout is reached",
					},
					cli.BoolFlag{
						Name:  "force, f",
						Usage: "Remove the volume forcefully if possible",
					},
				},
				ArgsUsage:   "[policy name]/[volume name]",
				Description: "Remove the volume for a policy, deleting its contents.",
				Usage:       "Remove a volume and its contents",
				Action:      VolumeRemove,
			},
			{
				Name:        "snapshot",
				Description: "Snapshot management tools",
				Usage:       "Snapshot management tools",
				Subcommands: []cli.Command{
					{
						Name:        "take",
						ArgsUsage:   "[policy name]/[volume name]",
						Description: "Take a snapshot for a volume now",
						Usage:       "Take a snapshot for a volume now",
						Action:      VolumeSnapshotTake,
					},
					{
						Name:        "list",
						ArgsUsage:   "[policy name]/[volume name]",
						Description: "List snapshots",
						Usage:       "List snapshots",
						Action:      VolumeSnapshotList,
					},
					{
						Name:        "copy",
						ArgsUsage:   "[policy name]/[volume name] [snapshot name] [new volume name]",
						Description: "Copies a volume with a given snapshot name to the new volume name. The policy will remain the same, as well as the volume parameters.",
						Usage:       "Copy a volume snapshot to a new volume",
						Action:      VolumeSnapshotCopy,
					},
				},
			},
			{
				Name:        "runtime",
				Description: "Runtime configuration management",
				Usage:       "Runtime configuration management",
				Subcommands: []cli.Command{
					{
						Name:        "get",
						ArgsUsage:   "[policy name]/[volume name]",
						Description: "Get runtime configuration",
						Usage:       "Get runtime configuration",
						Action:      VolumeRuntimeGet,
					},
					{
						Name:        "upload",
						ArgsUsage:   "[policy name]/[volume name]",
						Description: "Upload runtime configuration",
						Usage:       "Upload runtime configuration",
						Action:      VolumeRuntimeUpload,
					},
				},
			},
		},
	},
	{
		Name:  "use",
		Usage: "Manage Uses (hosts consuming resources)",
		Subcommands: []cli.Command{
			{
				Name:        "list",
				Usage:       "List uses",
				Description: "List the uses the apiserver knows about, in newline-delimited form.",
				ArgsUsage:   "",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "snapshots",
						Usage: "List snapshots instead of mounts",
					},
				},
				Action: UseList,
			},
			{
				Name:        "get",
				Usage:       "Get use info",
				Description: "Obtains the information on a specified use. Requires that you know the policy and image name.",
				ArgsUsage:   "[policy name]/[volume name]",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "snapshot",
						Usage: "Get the snapshot use instead of a mount",
					},
				},
				Action: UseGet,
			},
			{
				Name:        "force-remove",
				ArgsUsage:   "[policy name]/[volume name]",
				Usage:       "Forcefully remove use information",
				Description: "Force-remove a use. Use this to correct unuseing errors or failing hosts if necessary. Requires that you know the policy and image name.",
				Action:      UseTheForce,
			},
			{
				Name:        "exec",
				ArgsUsage:   "[policy name]/[volume name] -- [command] ...",
				Usage:       "Execute a command when use locks free.",
				Description: "When both snapshot and mount locks are freed, take the locks and execute the command.",
				Action:      UseExec,
			},
		},
	},
}
