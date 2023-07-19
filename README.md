# coffeeshop-clone
Clone of go-coffeeshop project to practice

Reference: https://github.com/thangchung/go-coffeeshop

## Specification
- Buf CLI: https://buf.build/docs/installation

Commands

```bash
$ buf mod init # init buf.yaml
$ buf mod update # generate buf.clock define deps
$ buf generate proto # generate code using plugin in buf.gen.yaml
```

## Project Structure

Reference: https://github.com/golang-standards/project-layout

### `/cmd`
- Main applications for this project

### `/internal`
- Private application and library code. This is the code you don't want others importing in their applications or libraries.
- Note that this layout pattern is enforced by the Go compiler itself (>= Go 1.4).

### `/pkg`
- Library code that's ok to use by external applications

### `/build`
- Packaging and Continuous Integration.

TBD