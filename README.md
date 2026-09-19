# LPE Checker [lpechk]
![GitHub release](https://img.shields.io/github/v/release/wxwreak/lpechk?style=for-the-badge&color=blue)
![GitHub all releases](https://img.shields.io/github/downloads/wxwreak/lpechk/total?style=for-the-badge&color=green)
![lpechk](https://github.com/wxwreak/lpechk/blob/main/lpechk.png)

---

A lightweight, automated local enumeration tool written in Go designed to audit Linux systems for potential privilege escalation vectors, misconfigurations, and sensitive data exposure.

---

## Features
- **User Content Analysis:** Analyzes the active session to detect current user privilages, UID, and GID parameters.
- **Sudo Permission Check:** Queries the system for passwordless or available sudo rights using non-interacative sudo validation.
- **SUID Binary Enum:** Recursively scans disk partitions while omitting pseudo-filesystems to flag dangerous or unexpected setuid executables.
- **Account Enum:** Parses local authentication files to list users configured with valid interactive command-line shells.
- **Credential Exposure Check:** Probes sensitive system file permissions to verify if restricted files like shadow storage are world-readable.
- **Environment Variable Audit:** Inspects runtime environment variables for insecure injection vectors, dangerous path configurations, and hardcoded secrets.

## Compilation
```bash
make all
```

## Usage
Run the compiled binary directly on the target system:
```bash
./dist/lpechk-linux-amd64
```

## Output Legend
- `[+]` **Informational status or succesfull checks.**
- `[!]` **Warnings or pontential points of interest.**
- `[!!!]` **Critical security findings or misconfigurations.**

# Disclaimer
This tool is created for educational purposes and authorized security testing only. The author is not responsible for any misuse of this software. Always ensure you have permission to test the target domain before running this tool.

## Contributing
Contributions are welcome! If you'd like to help improve **lpechk**, please read the [CONTRIBUTING.md](CONTRIBUTING.md) file to understand how to get started, report bugs, or submit your own pull requests.
