<p align="center">
	<img align="center" style="height: 30%;width:30%;" src="https://github.com/whiterabb17/Shamanic/blob/main/resources/Shaman.png">
</p>

# Shamanic BackDoor
〢public release ✅ <br>
〢private release contains far more features such as:
-  additional communication protocols 
- pivot/spread features
- recovery features

Shaman is a Resilient, Stealthy & capable Windows Backdoor.<br>
Utilizing Telegram for command relays, allowing on the go operations.<br><br>
Designed for a more stealthy & concrete foothold during red-team engagements.<br>
Agents gather in a `Hive` allowing both targeted and mass command taskings to all the connected agents.

Can be compiled for *Nix or Mac, as it's supported by the gryphon framework<br>
though this will require some manual tweaking to be done by the <b>user</b><br>
as Shaman uses certain windows api calls for antivm and some other stuff<br>

# NOTICE
To prevent misuse of more field ready builds by scriptkiddies
the obfuscated build will not compile without some manual tweaking by the user

# Latest scan of Obfuscated Build
<p align="center">
	<img align="center" style="height:50%;width:25%" src="https://antiscan.me/images/result/wKiQxZeWFduv.png">
	<img align="center" style="height:40%;width:60%" src="https://github.com/whiterabb17/Shamanic/blob/main/resources/vt.png">
</p>

# Required
Enter your:
- telegram userid 
- telegram botids 
- telegram bottokens

in <b>package/util/constants.go</b>

# Operation Requirements
2 Bot channels required.<br>
A `Library` channel which will hold client Identifiers and Heartbeats (online status)<br>
A `Dispatch` channel which will send and recieve commands and responses respectively

# Demo

<p align="center">
	<img align="center" style="height: 30%;width:30%;" src="https://github.com/whiterabb17/Shamanic/blob/main/resources/Shaman%20Demo.gif">
</p>

## Commands & Functions
	Base Commands & Functions
	===========================
	help 					Display this help message
	list 					Display currently open doors
	ping 					Measure the latency of command execution
	gryphon [command args] 			Execute Gryphon command w/out arguments <b>More info below</b>
	reset 					Create a new Summoning message
	info 					Display system information
	soft 					Display the list of installed programs
	sh 					Execute a command and return the output
	up 					Upload a file from the local system
	dl 					Download a file from a url to the local system
	root 					Ask for admin permissions
	inst 					Returns instance informtaion
	remove 					Uninstall Shaman bin & persistence
	
## Gryphon Offensive Module Functions
	
	Commands With Arguments
	===========================
	 SliceFile       [arg:  string]			Return slice from file
	 MakeZip         [args: string, []string]	Create $zip_name from $fileNames
	 DnsLookup       [arg:  string] 		Performs DNS Lookup of given hostname
	 RdnsLookup      [arg:  string] 		Performs reverse DNS Lookup of given IP
	 HostsPassive    [arg:  string] 		ARP Monitors networks at given interval
	 FilePermissions [arg:  string] 		Checks for read/write of given file
	 Portscan        [args: string, int] 		Performs multi-port scan
	 PortscanSingle  [args: string, int] 		Single port scan
	 BannerGrab      [args: string, int]
	 CmdOut          [arg:  string] 		Runs a cmd and returns output
	 CmdOutPlatform  [arg:  string] 		Platform aware cmd run with output return
	 CmdRun          [arg:  string] 		Runs cmd without return of data
	 CmdBlind        [arg:  string] 		Unsupervision cmd run, no output
	 CreateUser      [args: string, string] 	Temporarily on supported on windows
	 Bind            [arg:  int] 			Binds a shell to given port
	 Reverse         [args: string, int] 		Runs a reverse shell
	 SendDataTcp     [args: ip/host, int, string] 	Sends data to given host using TCP
	 SendDataUdp     [args: ip/host, int, string] 	Sends data to given host using UDP
	 ReadFile        [arg:  string]
	 WriteFile       [arg:  string]
	 IP2Hex          [arg:  string]
	 Port2Hex        [arg:  int]
	 Download        [arg:  string]
	 CopyFile        [arg:  string, string] 
	 PkillPid        [arg:  int]
	 PkillName       [arg:  string]
	 Persist         [arg:  string] 		Available options: Startup (Win/*Nix), Schtasks (Win ONLY)
	 SelfInject      [arg:  string] 		Url to download a file from to Inject bytes into owned Process
	 DropInject      [arg:  string] 		Url to download a file from to Inject after dropping on disk
	 ProcInject      [args: string, string]		Downloads binary from provided URL and injects into specified process
	 RefLoad         [arg:  string] 		Url to	download a file from to Reflectively Load into current domain
	
	
	Commands Without Arguments
	==============================
	 pkillAv        				Kills most common AV
	 clearLogs      				Clears most system logs
	 interfaces    					Gets network interfaces to use for Sniffing
	 sniffNetwork   				Starts a network traffic sniffer that writes traffic to file for retrieval
	 fetchNetLogs   				Retrieves Sniffer logs if they exist
	 listDir  					Returns files in yellow directory
	 networks       				Returns list of nearby Wi-Fi networks
	 localIP      					Gets Private IP
	 globalIP     					Gets Public IP
	 isroot         				Checks is client is running as root/admin
	 proc       					Returns processes and their PIDs
	 systeminfo     				Returns general system info
 	 escalate       				Attempts PrivilegeEscalation through various different methods
