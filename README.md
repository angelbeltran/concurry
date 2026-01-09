# Concurry

## CLI

This cli generates go types that wrap targeted types with the wrapper types having the same public methods but with all starting context.Context arguments curried.

### Installation

```
go install github.con/angelbeltran/concurry
```

**If using go 1.24+, to add as a tool dependency**
```
go get -tool github.con/angelbeltran/concurry
```
