goCfgMgr
========

A Json config manager written by Golang.

Usage
========

    go get github.com/reckhou/goCfgMgr

In code:

    import "github.com/reckhou/goCfgMgr"

The config is a json format text file including 3 sections by default:`basic`, `projects`, `additional`.

For example, if you want to get `Port` param under `basic` section, you might use this code:

    goCfgMgr.Get("basic", "Port").(string)

For more usage, just read the code, it's simple & elegant(for a noob gopher). :)
