package internal

const ColorRed = "\033[0;31m"
const ColorNone = "\033[0m"
const ColorGreen = "\033[0;32m"

const KeyFileLocation = "/usr/share/keyrings/kali.gpg"

const RepositoryFileLocation = "/etc/apt/sources.list.d/kali.list"
const RepositoryFileContent = "deb [signed-by=/usr/share/keyrings/kali.gpg] http://http.kali.org/kali kali-rolling main contrib non-free"
