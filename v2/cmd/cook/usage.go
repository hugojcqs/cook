package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/glitchedgitz/cook/v2/pkg/util"
)

func helpMode(h []string) {

	helpModeNames := func() string {
		t := ""
		for k := range helpFunctions {
			t += k + ", "
		}
		return t
	}()

	if len(h) <= 0 {
		log.Fatalf("Ask for these... %s", helpModeNames)
	}

	help := strings.ToLower(h[0])

	if fn, exists := helpFunctions[help]; exists {
		fn()
	} else {
		log.Fatalf("Ask for these... %s", helpModeNames)
	}

	os.Exit(0)
}

func showHelp() {
	fmt.Fprintln(os.Stderr, banner)
	fmt.Fprintln(os.Stderr, util.Reset)
	flagsHelp()
	os.Exit(0)
}

func printHelp(title string, description ...string) {
	fmt.Println(util.Blue + title + util.Reset)
	for _, d := range description {
		fmt.Println("    " + d)
	}
	fmt.Println()
}

func flagsHelp() {

	printHelp("GITHUB", "https://github.com/glitchedgitz/cook")

	printHelp(
		"USAGE",
		"cook [params and values] [pattern]",
		"cook -param1 value -file: filename -param3 value param3,admin_set param1 file,[1-300]",
	)

	printHelp(
		"MODES",
		"Search                     cook search [word]",
		"Help                       cook help [word]",
		"Update                     cook update [filename]",
		"                           This will update the file's cache.",
		"                              - Use \"cache\" to update cached file from source",
		"                              - Use \"db\" to update cooks-ingredients",
		"                              - Use \"*\" to do both",
		"Add                        cook add keyword=[values, separated by comma] in [category]",
		"                           	(files, raw-files, functions and lists)",
		"                               (This will only make changes in my.yaml)",
		"Delete                     cook delete [keyword]",
		"                               (This will only make changes in my.yaml)",
		"Show                       cook show [category]",
		"                           Better not try for \"files\"",
	)

	printHelp(
		"FLAGS",
		"        -peek              Peek the output using `-peek 50` for first 50 lines",
		"-a      -append            Append to the previous lines, instead of permutations",
		"-c      -col               Print column numbers and there values",
		"-conf,  -config            Config Information",
		"-mc,     -methodcol        Apply methods column wise",
		"                           	-mc 0:md5,b64e; 1:reverse",
		"                           To all cols separate",
		"                           	-mc md5,b64e",
		"-m,     -method            Apply methods to final output",
		"-h,     -help              Help",
		"        -min               Minimum no of columns to print",
		"-o,     -output            Write output to file instead of stdout",
		"                           	cook wordlist1 wordlist2 -o output.txt",
		"-ob,    -output-both       Write output to both file and stdout",
		"                           	cook wordlist1 wordlist2 -o output.txt -ob",
	)
}

func methHelp() {
	printHelp(
		"METHODS",
		"Apply different sets of operations to your wordlists",
	)
	printHelp(
		"STRING/LIST/JSON",
		"sort                           - Sort them",
		"sortu                          - Sort them with unique values only",
		"reverse                        - Reverse string",
		"split                          - split[char]",
		"splitindex                     - splitindex[char:index]",
		"replace                        - Replace All replace[this:tothis]",
		"leet                           - a->4, b->8, e->3 ...",
		"                                 leet[0] or leet[1]",
		"json                           - Extract JSON field",
		"                                 json[key] or json[key:subkey:sub-subkey]",
		"smart                          - Separate words with naming convensions",
		"                                 redirectUri, redirect_uri, redirect-uri  ->  [redirect, uri]",
		"smartjoin                      - This will split the words from naming convensions &",
		"                                 param.smartjoin[c,_] (case, join)",
		"                                 redirect-uri, redirectUri, redirect_uri ->  redirect_Uri",
		"",
		"u          upper               - Uppercase",
		"l          lower               - Lowercase",
		"t          title               - Titlecase",
	)
	printHelp(
		"URLS",
		"fb         filebase            - Extract filename from path or url",
		"s          scheme              - Extract http, https, gohper, ws, etc. from URL",
		"           user                - Extract username from url",
		"           pass                - Extract password from url",
		"h          host                - Extract host from url",
		"p          port                - Extract port from url",
		"ph         path                - Extract path from url",
		"f          fragment            - Extract fragment from url",
		"q          query               - Extract whole query from url",
		"k          keys                - Extract keys from url",
		"v          values              - Extract values from url",
		"d          domain              - Extract domain from url",
		"           tld                 - Extract tld from url",
		"           alldir              - Extract all dirrectories from url's path",
		"sub        subdomain           - Extract subdomain from url",
		"           allsubs             - Extract subdomain from url",
	)
	printHelp(
		"ENCODERS",
		"b64e       b64encode           - Base64 encoder",
		"hexe       hexencode           - Hex string encoder",
		"           charcode            - Give charcode encoding",
		"                                 charcode[0] without semicolon",
		"                                 charcode[1] with semicolon",
		"jsone      jsonescape          - JSON escape",
		"urle       urlencode           - URL encode reserved characters",
		"           utf16               - UTF-16 encoder (Little Endian)",
		"           utf16be             - UTF-16 encoder (Big Endian)",
		"xmle       xmlescape           - XML escape",
		"urleall    urlencodeall        - URL encode all characters",
		"unicodee   unicodeencodeall    - Unicode escape string encode (all characters)",
	)
	printHelp(
		"DECODERS",
		"b64d       b64decode           - Base64 decoder",
		"hexd       hexdecode           - Hex string decoder",
		"jsonu      jsonunescape        - JSON unescape",
		"unicoded   unicodedecode       - Unicode escape string decode",
		"urld       urldecode           - URL decode",
		"xmlu       xmlunescape         - XML unescape",
	)
	printHelp(
		"HASHES",
		"md5                            - MD5 sum",
		"sha1                           - SHA1 checksum",
		"sha224                         - SHA224 checksum",
		"sha256                         - SHA256 checksum",
		"sha384                         - SHA384 checksum",
		"sha512                         - SHA512 checksum",
	)
}

func usageHelp() {
	printHelp(
		"BASIC USAGE",
		"$ cook -start admin,root  -sep _,-  -end secret,critical  / start sep end",
		"$ cook / admin,root _,- secret,critical",
	)
	printHelp(
		"FUNCTIONS",
		"Use functions such as date for different variations of values",
		"$ cook -dob date(17,Sep,1994) elliot _,- dob",
	)
	printHelp(
		"RANGES",
		"Use ranges like 1-100, A-Z, a-z or A-z in pattern of command",
		"$ cook 1-999",
		"$ cook a-z",
		"$ cook A-Z",
		"$ cook X-d",
	)
	printHelp(
		"RAW STRINGS",
		"Print value without any parsing/modification.",
		"$ cook -date `date(17,Sep,1994)` date",
	)
	printHelp(
		"PIPE INPUT",
		"Use - as param value for pipe input",
		"$ cook -d - d / test",
	)
	printHelp(
		"USING -min",
		"Print value without any parsing/modification",
		"$ cook n n n -min 1",
	)
}
