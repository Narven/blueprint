# blueprint

Blueprint is a tool for generating a specific folders/files structure base on a predefined template.

Basic
```bash
blueprint simple
```

This will generate a folder/file structure based on the contents of the `simple.yml` located in the
`~/.blueprint`

Examples:

```yaml
version: 1

name: Simple Blueprint

output: ~

structure:
  - FolderA/
  - FolderB/
  - FolderC/
  - FolderD/
  - simple
```

this would generate the following structure in your home folder

```
~/
 FolderA
 FolderB
 FolderC
 FolderD
 simple
```