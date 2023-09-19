// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["turbo"] = model.Subcommand{
		Name:        []string{"turbo"},
		Description: `Turborepo is a high-performance build system for JavaScript and TypeScript codebases`,
		Options: []model.Option{{
			Name:        []string{"--version"},
			Description: `Print the version`,
		}, {
			Name:         []string{"--help"},
			Description:  `Print a help message`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"bin"},
			Description: `Get the path to the turbo binary`,
		}, {
			Name:        []string{"link"},
			Description: `Link your local directory to a Vercel organization and enable remote caching`,
			Options: []model.Option{{
				Name:        []string{"--no-gitignore"},
				Description: `Do not create or modify .gitignore`,
			}},
		}, {
			Name:        []string{"login"},
			Description: `Login to your Vercel account`,
			Options: []model.Option{{
				Name:        []string{"--sso-team"},
				Description: `Attempt to authenticate to the specified team using SSO`,
			}},
		}, {
			Name:        []string{"logout"},
			Description: `Logout of your Vercel account`,
		}, {
			Name:        []string{"prune"},
			Description: `Prepare a subset of your monorepo`,
			Options: []model.Option{{
				Name:        []string{"--scope"},
				Description: `Specify package to act as entry point for pruned monorepo`,
				Args: []model.Arg{{
					Name: "package",
				}},
			}, {
				Name:        []string{"--docker"},
				Description: `Output pruned workspace into 'fill' and 'json' directories optimized for Docker layer caching`,
			}},
		}, {
			Name:        []string{"unlink"},
			Description: `Unlink the current directory from your Vercel organization and disable remote caching`,
		}, {
			Name:        []string{"run"},
			Description: `Run tasks in your monorepo`,
			Args: []model.Arg{{
				Name:       "tasks",
				Generator:  nil, // TODO: port over generator
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"--scope"},
				Description: `Specify packages to act as entry points for task execution`,
				Args: []model.Arg{{
					Name: "package",
				}},
			}, {
				Name:        []string{"--cache-dir"},
				Description: `Specify local filesystem cache directory`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFolders},
					Name:      "dir",
				}},
			}, {
				Name:        []string{"--concurrency"},
				Description: `Limit the concurrency of task execution (use "1" for serial)`,
				Args: []model.Arg{{
					Name: "limit",
				}},
			}, {
				Name:        []string{"--continue"},
				Description: `Continue execution even if a task exits with an error`,
			}, {
				Name:        []string{"--force"},
				Description: `Ignore the existing cache`,
			}, {
				Name:        []string{"--graph"},
				Description: `Generate a Dot graph of the task execution`,
			}, {
				Name:        []string{"--global-deps"},
				Description: `Specify glob of global filesystem dependencies to be hashed (useful for .env and files in the root directory)`,
				Args: []model.Arg{{
					Name: "glob",
				}},
			}, {
				Name:        []string{"--since"},
				Description: `Limit/set scope to changed packages since a mergebase`,
				Args: []model.Arg{{
					Name:      "branch",
					Generator: nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"--team"},
				Description: `The slug of a turborepo.com team`,
				Args: []model.Arg{{
					Name: "slug",
				}},
			}, {
				Name:        []string{"--token"},
				Description: `A turborepo.com access token`,
				Args: []model.Arg{{
					Name: "token",
				}},
			}, {
				Name:        []string{"--ignore"},
				Description: `Files to ignore when calculating changed files (supports globs)`,
				Args: []model.Arg{{
					Name: "files",
				}},
			}, {
				Name:        []string{"--profile"},
				Description: `File to write turbo's performance profile into`,
				Args: []model.Arg{{
					Templates: []model.Template{model.TemplateFilepaths},
					Name:      "file",
				}},
			}, {
				Name:        []string{"--parallel"},
				Description: `Execute all tasks in parallel`,
			}, {
				Name:        []string{"--include-dependencies"},
				Description: `Include the dependencies of tasks in execution`,
			}, {
				Name:        []string{"--no-deps"},
				Description: `Exclude dependent task consumers from execution`,
			}, {
				Name:        []string{"--no-cache"},
				Description: `Avoid saving task results to the cache (useful for development/watch tasks)`,
			}, {
				Name: []string{"--output-logs"},
				Args: []model.Arg{{
					Name: "level",
					Suggestions: []model.Suggestion{{
						Name:        []string{`new-only`},
						Description: `Only new output with hashes for cached tasks`,
					}, {
						Name:        []string{`hash-only`},
						Description: `Only turbo-computed task hashes`,
					}, {
						Name:        []string{`full`},
						Description: `Show all output`,
					}, {
						Name:        []string{`none`},
						Description: `Hide process output`,
					}},
				}},
			}, {
				Name:        []string{"--dry", "--dry-run"},
				Description: `List the packages in scope and the tasks that would be run`,
				Args: []model.Arg{{
					Name:        "format",
					Suggestions: []model.Suggestion{{Name: []string{`json`}}},
					IsOptional:  true,
				}},
			}},
		}},
	}
}