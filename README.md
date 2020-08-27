# vfsgen-example

My example of using [vfsgen](https://github.com/shurcooL/vfsgen) with modules.

Avoids using `vfsgendev` or a "generate comment"   

Run dev, read files from disk

    go run -tags dev dev.go
    
Run with Virtual File System

    # Delete generated
    ./reset.sh
    
    # Generate VFS
    go run -tags vfs static/ui_vfsgen.go
    
    # Read files from static/ui/vfs.go 
    go run main.go
    
    