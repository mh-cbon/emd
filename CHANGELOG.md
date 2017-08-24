# Changelog - emd

### 1.0.0

__Changes__

- toc: fix link generator to remove colon
- close #26: avoid fatal error when a path is not recognized
- appveyor: close #28: fix badge urls

__Contributors__

- mh-cbon
- solvingJ

Released by mh-cbon, Thu 24 Aug 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.12...1.0.0#diff)
______________

### 0.0.12

__Changes__

- __cli__:
  - ensure data passed on the command line overwrites everthing else, fix #19.

- __dep__:
  - fixed glide lock, needed an update.

- __dev__:
  - Add provider tests





__Contributors__

- mh-cbon

Released by mh-cbon, Wed 10 May 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.11...0.0.12#diff)
______________

### 0.0.11

__Changes__

- __CLI__:
  - Fix an issue in argument handling.

- __Path detection__: #17
  - The mechanism to detect and extract information from the path,
    is improved to work with anything that matches src/[provider]/[user]/[repo]/[path...]

- __Temaplate func helpers__:
  - __pkg_doc__: won t panic anymore if the default `main.go` file is not yet created.

- __Temaplates__: #16
  - Fix templates output to avoid double `/`.

- __dev__:
  - Largely improved tests.







__Contributors__

- mh-cbon
- suntong

Released by mh-cbon, Sun 07 May 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.10...0.0.11#diff)
______________

### 0.0.10

__Changes__

- __CLI__:
  - new `init` command: Writes a basic `README.e.md` frile to get started
  - `gen`: removed useless confirmation message on successful operation.
  - `stdin`: emd can now receives the input template from STDIN.
  - verbosity: Added support for verbosity with env variable `VERBOSE=y`.

- __Predefined data__:
  - New api is introduced to better detect predefined data.
  - ProjectPath is better handled when its a symlink outside of `GOPATH` (#15 thanks suntong)

- __Prelude data__:
  - It is now possible to override predefined data within the prelude.
    This should allow the end user to recover from a buggy implementation in the pre defined variables declaration.

- __TOC__:
  - #16 Improve link generator to handle `[]|',`
  - Multiple fixes to properly render the TOC level of each header.

- __Template__: added multiple new functions to help to work with templates
  - __set__(name string, x interface{}): Save given value `x` as `name` on dot `.`.
  - __link__(url string, text ...string) string: Prints markdown link.
  - __img__(url string, alt ...string) string: Prints markdown image.
  - __concat__(x ...string) string: Concat given arguments.
  - __pathjoin__(x ...string) string: Join given arguments with `/`.

- __dev__:
  - added small test suites into `test.sh`









__Contributors__

- mh-cbon
- suntong

Released by mh-cbon, Sat 06 May 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.10-beta...0.0.10#diff)
______________

### 0.0.10-beta

__Changes__

- CLI:
  - new `init` command: Writes a basic `README.e.md` frile to get started
  - `gen`: removed useless confirmation message on successful operation.
- Predefined data:
  - New api is introduced to better detect predefined data.
- Prelude data:
  - It is now possible to override predefined data within the prelude.
    This should allow the end user to recover from a buggy implementation in the pre defined variables declaration.
- TOC:
  - #16 Improve link generator to handle `[]|',`
  - Multiple fixes to properly render the TOC level of each header.
- Template: added multiple new functions to help to work with templates
  - __set__(name string, x interface{}): Save given value `x` as `name` on dot `.`.
  - __link__(url string, text ...string) string: Prints markdown link.
  - __img__(url string, alt ...string) string: Prints markdown image.
  - __concat__(x ...string) string: Concat given arguments.
  - __pathjoin__(x ...string) string: Join given arguments with `/`.











__Contributors__

- mh-cbon

Released by mh-cbon, Mon 24 Apr 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.9...0.0.10-beta#diff)
______________

### 0.0.9

__Changes__

- feature #10: Emd file can define a prelude block of `yaml` data to inject into the template processing.
- cli fix: Before the -in argument was mandatory. It was not possible to use the default template.
- feature installer: Fixed apt/rpm repositories.
- feature #8: ensure an errored command line execution displays correctly.
- feature: added new template functions
  - __yaml__(file string, keypaths ...string): parses and build new yaml content of given file.
  - __preline__(pre , content string): prepends pre for every lines of content.
  - __echo__(s ...string): echo every string s.
  - __read__(file string): returns file content.
  - __cat__(file string): to display the file content.
  - __exec__(bin string, args ...string): to exec a program.
  - __shell__(s string): to exec a command line on the underlying shell (it is not cross compatible).
  - __color__(color string, content string): to embed content in a block code with color.
  - __gotest__(rpkg string, run string, args ...string): exec `go test <rpkg> -v -run <run> <args...>`.
  - __toc__(maximportance string, title string): display a TOC.
- feature: added new badge templates
  - __license/shields__: show a license badge
  - __badge/codeship__: show a codeship badge
- __deprecation__ #7: some template functions were deprecated to avoid pre defined formatting,
      old behavior can be reset via new options defined into the prelude data. See also the new __color__ function.
  - __file__ is dprecated for __cat__
  - __cli__ is dprecated for __exec__
- badges fix #14: removed useless whitespace
- dev: Added support for glide, it was required to handle yaml.
- dev: updated tests
- dev: godoc documentation improvements.

__Contributors__

- suntong
- mh-cbon

Released by mh-cbon, Sat 22 Apr 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.9-beta9...0.0.9#diff)
______________

### 0.0.9-beta9

__Changes__

- fix default reading of the md file

__Contributors__

- mh-cbon

Released by mh-cbon, Tue 18 Apr 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.9-beta8...0.0.9-beta9#diff)
______________

### 0.0.9-beta8

__Changes__

- fix #15: properly handle symbolic links
- fix cli: it was not possible to use the default template

__Contributors__

- mh-cbon

Released by mh-cbon, Tue 18 Apr 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.9-beta7...0.0.9-beta8#diff)
______________

### 0.0.9-beta7

__Changes__

- fix wrong import path

__Contributors__

- mh-cbon

Released by mh-cbon, Mon 17 Apr 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.9-beta6...0.0.9-beta7#diff)
______________

### 0.0.9-beta6

__Changes__

- ci: fix scripts to add glide support

__Contributors__

- mh-cbon

Released by mh-cbon, Mon 17 Apr 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.9-beta5...0.0.9-beta6#diff)
______________

### 0.0.9-beta5

__Changes__

- new functions:
  - __yaml__(file string, keypaths ...string): parses and build new yaml content of given file.
  - __preline__(pre , content string): prepends pre for every lines of content.
  - __echo__(s ...string): echo every string s.
  - __read__(file string): returns file content.
- toc: multiple fixes,
    - it properly handles duplicated title by appending an increment
    - fix handling of !; in links generator
    - fix line counting when extracting markdown titles
    - fix md title selection starting at line N
- prelude:
  - fix read of quoted values
  - prelude data is now read on all templates added, not only file
  - fix last eol handling
- codeship fix #2: proper project url
- bump script: added new utils/ tests
- glide: init versionned dependencies to handle yaml files.
- godoc: refactoring to improve documentation

__Contributors__

- mh-cbon

Released by mh-cbon, Mon 17 Apr 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.9-beta4...0.0.9-beta5#diff)
______________

### 0.0.9-beta4

__Changes__

- toc: improve toc parser, refactored, added tests

__Contributors__

- mh-cbon

Released by mh-cbon, Fri 14 Apr 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.9-beta3...0.0.9-beta4#diff)
______________

### 0.0.9-beta3

__Changes__

- fix #13: add new template to show a license badge.
- prelude: trim leading whitespaces of unquoted values.
- fix #14: improved badge output, removed useless whitespace.
- fix #2: codeship badge template, added a CsProjectID parameter.
- exec/shell/cat/gotest: avoid pre defined formatting, old behavior can be reset via new options defined into the prelude data.
- toc: fixed some corner cases while parsing/generating the TOC.

__Contributors__

- mh-cbon

Released by mh-cbon, Fri 14 Apr 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.9-beta2...0.0.9-beta3#diff)
______________

### 0.0.9-beta2

__Changes__

- fix some bugs in TOC title evaluation and generation
- fix apt repository!

__Contributors__

- mh-cbon

Released by mh-cbon, Thu 13 Apr 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.9-beta1...0.0.9-beta2#diff)
______________

### 0.0.9-beta1

__Changes__

- deprecation: improve error messages

__Contributors__

- mh-cbon

Released by mh-cbon, Wed 12 Apr 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.9-beta...0.0.9-beta1#diff)
______________

### 0.0.9-beta

__Changes__

- close #10: added feature to read, decode and registers the prelude data

  It is now possible to define a prelude block of `yaml` data in your __README__ file to
  register new data.

- added __cat/exec/shell/color/gotest/toc__ func

  - __cat__(file string): to display the file content.
  - __exec__(bin string, args ...string): to exec a program.
  - __shell__(s string): to exec a command line on the underlying shell (it is not cross compatible).
  - __color__(color string, content string): to embed content in a block code with color.
  - __gotest__(rpkg string, run string, args ...string): exec `go test <rpkg> -v -run <run> <args...>`.
  - __toc__(maximportance string, title string): display a TOC.

- close #7: deprecated __file/cli__ func

  Those two functions are deprecated in flavor of their new equivalents,
  __cat/exec__.

  The new functions does not returns a triple backquuotes block code.
  They returns the response body only.
  A new function helper __color__ is a added to create a block code.

- close #8: improved cli error output

  Before the output error was not displaying
  the command line entirely when it was too long.
  Now the error is updated to always display the command line with full length.

- close #9: add new gotest helper func
- close #12: add toc func
- close #10: ensure unquoted strings are read properly
- close #11: add shell func helper.

__Contributors__

- mh-cbon

Released by mh-cbon, Wed 12 Apr 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.8...0.0.9-beta#diff)
______________

### 0.0.8

__Changes__

- fix goreport badge template

__Contributors__

- mh-cbon

Released by mh-cbon, Sun 12 Mar 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.7...0.0.8#diff)
______________

### 0.0.7

__Changes__

- improve template documentation
- goreport: add template (fixes #4)

__Contributors__

- mh-cbon

Released by mh-cbon, Sun 12 Mar 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.6...0.0.7#diff)
______________

### 0.0.6

__Changes__

- template functions (std): add a new render template function to define additional values (fixes #2)
- template function (std): file takes a new argument to define the colorizer (fixes #1)
- emd: add new methods to access template, out and data
- release: fix missing version to the emd build
- README: multiple improvements.

__Contributors__

- mh-cbon

Released by mh-cbon, Mon 06 Mar 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.5...0.0.6#diff)
______________

### 0.0.5

__Changes__

- badges: add codeship
- Funcs cli/file: changed the MD template to add support for html anchors (before they was using bold tag, now they use a title tag)
- command gen: prints success message only if out is not stdout
- README: added a section to show HTML generation, and a recipe to bump the package.
- release: change bump script format

__Contributors__

- mh-cbon

Released by mh-cbon, Mon 06 Mar 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.4...0.0.5#diff)
______________

### 0.0.4

__Changes__

- changelog: typos
- README: add template helpers documentation

__Contributors__

- mh-cbon

Released by mh-cbon, Wed 22 Feb 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.3...0.0.4#diff)
______________

### 0.0.3

__Changes__

- travis(token): update ghtoken

__Contributors__

- mh-cbon

Released by mh-cbon, Wed 22 Feb 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.2...0.0.3#diff)
______________

### 0.0.2

__Changes__

- README: fix appveyor badge
- badge(update): fix url
- README: fix appveyor badge
- badge(fix): fix appveyor badge
- README: add appveyor badge
- badge(update): update text displayed in ci badges
- README(fix): use correct bin path
- bump(fix): emd gen command was wrong

__Contributors__

- mh-cbon

Released by mh-cbon, Wed 22 Feb 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/0.0.1...0.0.2#diff)
______________

### 0.0.1

__Changes__

- project initialization

__Contributors__

- mh-cbon

Released by mh-cbon, Wed 22 Feb 2017 -
[see the diff](https://github.com/mh-cbon/emd/compare/9b73c280847b824e4e366bcf3276d4eefecde4de...0.0.1#diff)
______________


