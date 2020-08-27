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
    
Modify `static/ui/data/docs/index.html` while running `dev.go` and `main.go`.
The former will print the current contents of the file,
the latter will print the file as it was when the VFS was generated 
    