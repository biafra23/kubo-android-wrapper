# Kubo wrapper to run Kubo on Android

## Build the wrapper library

```bash
go mod tidy 
go get golang.org/x/mobile/bind 
go get google.golang.org/genproto@latest
gomobile bind -ldflags "-checklinkname=0" -target=android -o wrapper.aar ./wrapper
```
## Add the wrapper library to your Android project
1. Copy the `wrapper.aar` file to your Android project's `libs` directory.
2. Add the following lines to your `build.gradle` file:

```Kotlin
dependencies {
    implementation(fileTree(mapOf("dir" to "libs", "include" to listOf("*.aar"))))
}
```

## Usage

```Kotlin
val ipfsRepoPath = applicationContext.filesDir.absolutePath + "/ipfs_repo"
// start Kubo as a daemon and initialize repo if necessary
Wrapper.runCli(ipfsRepoPath, "DEBUG", "daemon --init")
```

## Notes

- The wrapper library is built using gomobile bind, which generates a Java/Kotlin binding for the Go code.
- The `runCli` method is used to start the Kubo daemon with the specified log level and command-line arguments.
- The `ipfsRepoPath` is the path to the IPFS repository, which is created in the app's private storage directory.
- The `DEBUG` log level is used for debugging purposes. You can change it to `INFO`, `WARN`, or `ERROR` as needed.
- The `daemon --init` command initializes the IPFS repository if it does not exist. You can change this to other commands as needed.
- The `runCli` method is a blocking call, meaning it will not return until the Kubo daemon is stopped. You may want to run it in a separate thread or use coroutines to avoid blocking the main thread.
- Make sure to handle the permissions for accessing the file system in your Android app, especially if you are targeting Android 10 or higher. You may need to request permissions at runtime.

