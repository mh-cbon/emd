# Changelog - emd

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


