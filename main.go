package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

func main() {
	var dangerSUID = map[string]bool{
		"aa-exec": true, "ab": true, "alpine": true, "ansible-test": true, "aout-read": true,
		"apache2": true, "aplay": true, "apropos": true, "apt-get": true, "apt": true,
		"ar": true, "aria2c": true, "arj": true, "as": true, "ascii-xfr": true,
		"ascii85": true, "ash": true, "aspell": true, "at": true, "atobm": true,
		"awk": true, "aws": true, "base32": true, "base58": true, "base64": true,
		"basenc": true, "bash": true, "bat": true, "bc": true, "bbyy": true,
		"bzip2": true, "c89": true, "c99": true, "cabal": true, "cancel": true,
		"capsh": true, "cat": true, "certbot": true, "checkv": true, "chgrp": true,
		"chmod": true, "chown": true, "chroot": true, "cisco": true, "clamscan": true,
		"cmp": true, "column": true, "comm": true, "composer": true, "cowdict": true,
		"cowsay": true, "cp": true, "cpan": true, "cpio": true, "cpulimit": true,
		"crash": true, "crontab": true, "csh": true, "csplit": true, "csvtool": true,
		"cupsfilter": true, "curl": true, "cut": true, "dash": true, "date": true,
		"dbadm": true, "dbx": true, "dc": true, "dd": true, "debugfs": true,
		"dialog": true, "diff": true, "dig": true, "din": true, "dir": true,
		"dmidecode": true, "dmsetup": true, "dnf": true, "docker": true, "dosbox": true,
		"dot": true, "dpkg": true, "dread": true, "dsh": true, "dttw": true,
		"du": true, "dvips": true, "easy_install": true, "eb": true, "echo": true,
		"ed": true, "efax": true, "emacs": true, "enjarify": true, "env": true,
		"eqn": true, "es": true, "ex": true, "exiftool": true, "expand": true,
		"expect": true, "expr": true, "facter": true, "fattach": true, "fbi": true,
		"fc": true, "fetch": true, "ffmpeg": true, "find": true, "finger": true,
		"fish": true, "flock": true, "fmt": true, "fold": true, "fping": true,
		"ftp": true, "function": true, "gawk": true, "gcc": true, "gdb": true,
		"gem": true, "genie": true, "genisoimage": true, "getfacl": true, "ghc": true,
		"ghci": true, "gimp": true, "git": true, "gjs": true, "gkill": true,
		"gmatrix": true, "gmic": true, "gmocha": true, "gnat": true, "gnatmake": true,
		"gnd": true, "gnm": true, "gnuplot": true, "go": true, "gofmt": true,
		"gordon": true, "gpasswd": true, "gpick": true, "gpg": true, "gprof": true,
		"grep": true, "grub": true, "gsl": true, "gsql": true, "gtar": true,
		"gzip": true, "hd": true, "head": true, "hexdump": true, "highlight": true,
		"hostid": true, "hostname": true, "hpage": true, "hping3": true, "html2ps": true,
		"htop": true, "huntsman": true, "iconv": true, "id": true, "ifconfig": true,
		"igm": true, "image": true, "import": true, "install": true, "ionice": true,
		"ip": true, "irb": true, "ispell": true, "jjs": true, "join": true,
		"journalctl": true, "jq": true, "jrunscript": true, "jse": true, "jshell": true,
		"juicer": true, "kbd_mode": true, "kernel": true, "keychain": true, "kill": true,
		"ksh": true, "kss": true, "last": true, "latex": true, "ld.so": true,
		"ldconfig": true, "ldd": true, "less": true, "lex": true, "lftp": true,
		"lg": true, "libs": true, "line": true, "link": true, "ln": true,
		"local": true, "locate": true, "logread": true, "logsave": true, "look": true,
		"lorder": true, "lp": true, "lpc": true, "lpr": true, "ls": true,
		"lsattr": true, "lsb": true, "lsof": true, "lspci": true, "ltrace": true,
		"lua": true, "lualatex": true, "luatex": true, "lwp-download": true, "lwp-request": true,
		"lx": true, "lynx": true, "lz4": true, "lzma": true, "m4": true,
		"mail": true, "make": true, "man": true, "map": true, "mawk": true,
		"mcs": true, "md5sum": true, "merge": true, "mesa": true, "mesg": true,
		"messages": true, "mime": true, "mira": true, "mkdir": true, "mkfifo": true,
		"mknod": true, "mktemp": true, "mlocate": true, "mm": true, "mo": true,
		"mocha": true, "modprobe": true, "more": true, "mos": true, "mount": true,
		"mpv": true, "msgattrib": true, "msgcat": true, "msgconv": true, "msgfilter": true,
		"msgfmt": true, "msggrep": true, "msginit": true, "msgmerge": true, "msguniq": true,
		"mtr": true, "mv": true, "mysql": true, "nano": true, "nasm": true,
		"nawk": true, "nc": true, "ncat": true, "ne": true, "netcat": true,
		"netstat": true, "nice": true, "nl": true, "nm": true, "nmap": true,
		"node": true, "nohup": true, "nping": true, "npm": true, "nroff": true,
		"nsenter": true, "nv": true, "od": true, "of": true, "og": true,
		"openssl": true, "openvt": true, "opkg": true, "os": true, "otp": true,
		"package": true, "pager": true, "pandoc": true, "parted": true, "pascal": true,
		"passwd": true, "paste": true, "patch": true, "pax": true, "pdb": true,
		"pdflatex": true, "pdftex": true, "perf": true, "perl": true, "perror": true,
		"pf": true, "pg": true, "php": true, "pic": true, "pico": true,
		"pidstat": true, "pigz": true, "ping": true, "pip": true, "pkexec": true,
		"pkg": true, "pkg_info": true, "pkginfo": true, "pld": true, "plenv": true,
		"plot": true, "pman": true, "pmie": true, "pnm": true, "pod2man": true,
		"pod2text": true, "poly": true, "pop": true, "port": true, "post": true,
		"pr": true, "printenv": true, "printf": true, "proc": true, "prof": true,
		"prolog": true, "ps": true, "psftp": true, "psql": true, "ptx": true,
		"puppet": true, "pure": true, "push": true, "pwd": true, "py": true,
		"pydoc": true, "pyenv": true, "python": true, "qalk": true, "qemu": true,
		"rake": true, "ranlib": true, "raw": true, "rc": true, "rdesktop": true,
		"readelf": true, "readlink": true, "red": true, "redcarpet": true, "redis": true,
		"rev": true, "rex": true, "rf": true, "rlogin": true, "rlwrap": true,
		"rm": true, "rmdir": true, "Rscript": true, "rsync": true, "rt": true,
		"ruby": true, "run-mailcap": true, "run-parts": true, "rview": true, "rvim": true,
		"sash": true, "sbcl": true, "sc": true, "scalac": true, "scapy": true,
		"scp": true, "screen": true, "script": true, "scrot": true, "sd": true,
		"sed": true, "see": true, "service": true, "setarch": true, "setfacl": true,
		"setlock": true, "setsid": true, "sftp": true, "sh": true, "sha1sum": true,
		"sha224sum": true, "sha256sum": true, "sha384sum": true, "sha512sum": true, "shred": true,
		"shuf": true, "shutdown": true, "singularity": true, "size": true, "slsh": true,
		"smbget": true, "snow": true, "socat": true, "soelim": true, "sort": true,
		"source": true, "spec": true, "split": true, "sqlite3": true, "ss": true,
		"ssh-agent": true, "ssh-keygen": true, "ssh-keyscan": true, "ssh": true, "sshpass": true,
		"start-stop-daemon": true, "stat": true, "strace": true, "strings": true, "su": true,
		"sudo": true, "sum": true, "sysctl": true, "systemctl": true, "systemd-resolve": true,
		"tac": true, "tail": true, "tar": true, "task": true, "taskset": true,
		"tbl": true, "tclsh": true, "tcpdump": true, "tee": true, "telnet": true,
		"tex": true, "tftp": true, "tic": true, "time": true, "timedatectl": true,
		"timeout": true, "tmux": true, "top": true, "touch": true, "tpage": true,
		"tr": true, "tracepath": true, "traceroute": true, "troff": true, "tset": true,
		"tsort": true, "tty": true, "ubuntu": true, "ul": true, "unexpand": true,
		"uniq": true, "unix2dos": true, "unlz4": true, "unlzma": true, "unshare": true,
		"unxz": true, "unzip": true, "update-alternatives": true, "uptime": true, "url": true,
		"userhelper": true, "uuencode": true, "uudecode": true, "vboxmanage": true, "vi": true,
		"view": true, "vigr": true, "vim": true, "viman": true, "vipw": true,
		"virsh": true, "volatility": true, "w": true, "wall": true, "watch": true,
		"wc": true, "wget": true, "whiptail": true, "who": true, "whoami": true,
		"wic": true, "wild": true, "wish": true, "write": true, "xargs": true,
		"xdotool": true, "xelatex": true, "xeteX": true, "xmodmap": true, "xmore": true,
		"xox": true, "xxd": true, "xz": true, "yarn": true, "yash": true,
		"yp": true, "z": true, "zcalc": true, "zenity": true, "zip": true,
		"zless": true, "zmore": true, "zsh": true, "zstat": true, "zypper": true,
	}

	sysVectors := map[string]bool{
		// Basic libs
		"PATH":              true,
		"LD_PRELOAD":        true,
		"LD_LIBRARY_PATH":   true,
		"PYTHONPATH":        true,
		"LD_AUDIT":          true,
		"LD_CONFIG_TRACES":  true,
		"LD_DEBUG":          true,
		"LD_DEBUG_OUTPUT":   true,
		"LD_PROFILE":        true,
		"LD_PROFILE_OUTPUT": true,
		"LD_USE_LOAD_BIAS":  true,
		// Lang runtimes
		"PERL5OPT":     true,
		"PERL5LIB":     true,
		"RUBYOPT":      true,
		"RUBYLIB":      true,
		"NODE_OPTIONS": true,
		// Glibc
		"GCONV_PATH":     true,
		"GLIBC_TUNABLES": true,
		// Sudoedit
		"EDITOR": true,
		"VISUAL": true,
		"PAGER":  true,
		// Shell & Sys Vectors
		"IFS":            true,
		"BASH_ENV":       true,
		"ENV":            true,
		"SHELLOPTS":      true,
		"BASHOPTS":       true,
		"PROMPT_COMMAND": true,
	}
	keyWords := []string{"pass", "key", "secret", "token", "appi", "auth", "admin"}
	red := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	const banner = `
   __   ___  ____  _______           __          
  / /  / _ \/ __/ / ___/ /  ___ ____/ /_____ ____
 / /__/ ___/ _/  / /__/ _ \/ -_) __/  '_/ -_) __/		github.com/wxwreak
/____/_/  /___/  \___/_//_/\__/\__/_/\_\\__/_/     
	`
	fmt.Println(banner)
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("%s Error while detecting user: %s\n", yellow("[!]"), err)
	} else {
		fmt.Printf("%s Current User: %s [UID: %s ; GID %s]\n", cyan("[+]"), currentUser.Username, currentUser.Uid, currentUser.Gid)
		if currentUser.Uid == "0" {
			fmt.Printf("\n%s You are running as root\n", yellow("[!]"))
		}
	}
	// Sudo permissions detection
	cmd := exec.Command("sudo", "-n", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s Sudo is unavailable or requires a passwd\n", green("[+]"))
	} else {
		fmt.Printf("%s Sudo perms detected\n", cyan("[+]"))
		fmt.Println(string(output))
	}
	// Searching for Binaries
	ignorePaths := map[string]bool{
		"/proc":  true,
		"/sys":   true,
		"/dev":   true,
		"/run":   true,
		"/snap":  true,
		"/mnt":   true,
		"/media": true,
	}
	searchDir := "/"
	filepath.Walk(searchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			if ignorePaths[path] {
				return filepath.SkipDir
			}
			return nil
		}
		if sysStat, ok := info.Sys().(*syscall.Stat_t); ok {
			if sysStat.Mode&syscall.S_ISUID != 0 {
				binaryNm := filepath.Base(path)
				if dangerSUID[binaryNm] {
					fmt.Printf("%s Critical SUID: %s (Owner UID: %d)\n", red("[!!!]"), path, sysStat.Uid)
				} else {
					fmt.Printf("%s SUID Found: %s (Owner UID: %d)\n", yellow("[!]"), path, sysStat.Uid)
				}
			}
		}
		return nil
	})
	// Testing /etc/passwd
	file, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Printf("%s Can't open /etc/passwd\n", green("[+]"))
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, ":")
			if len(parts) < 7 {
				continue
			}
			shell := parts[6]
			if strings.HasSuffix(shell, "nologin") || strings.HasSuffix(shell, "false") {
				continue
			}
			if strings.Contains(line, "bash") || strings.Contains(line, "sh") || strings.Contains(line, "zsh") {
				fmt.Printf("%s User with shell: %s\n", cyan("[+]"), line)
			}
		}
		defer file.Close()
	}
	// Testing /etc/shadow
	file, err = os.Open("/etc/shadow")
	if err != nil {
		fmt.Printf("%s Can't open /etc/shadow\n", green("[+]"))
	} else {
		defer file.Close()
		fmt.Printf("%s Succesfully opened /etc/shadow\n", red("[!!!]"))
	}
	// Searching for Keywords & System Vectors
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) < 2 || parts[1] == "" {
			continue
		}
		key, value := parts[0], parts[1]
		lowrKey := strings.ToLower(key)

		if sysVectors[key] {
			fmt.Printf("%s System Vector: %s = %s\n", yellow("[!]"), key, value)

			if key == "PATH" {
				if strings.HasPrefix(value, ":") || strings.HasSuffix(value, ":") || strings.Contains(value, "::") {
					fmt.Printf("%s PATH contains an empty path\n", red("[!!!]"))
				}

				paths := filepath.SplitList(value)
				for _, p := range paths {
					if p == "." {
						fmt.Printf("%s Dot ('.') Found in PATH\n", red("[!!!]"))
					}
					if err := syscall.Access(p, 2); err == nil {
						fmt.Printf("%s Writable PATH directory found: %s\n", red("[!!!]"), p)
					}
				}
			}
			continue
		}
		for _, keyword := range keyWords {
			if strings.Contains(lowrKey, keyword) {
				fmt.Printf("%s Sensitive Data [%s]: %s = %s\n", yellow("[!]"), keyword, key, value)
				break
			}
		}
	}

	fmt.Printf("%s Done.\n", green("[+]"))
}
