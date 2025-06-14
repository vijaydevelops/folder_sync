package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"

    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    // "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2"
)

func syncFolders(src, dst string) error {
    return filepath.Walk(src, func(srcPath string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        relPath, err := filepath.Rel(src, srcPath)
        if err != nil {
            return err
        }

        dstPath := filepath.Join(dst, relPath)

        if info.IsDir() {
            return os.MkdirAll(dstPath, info.Mode())
        }

        dstInfo, err := os.Stat(dstPath)
        if err != nil || info.ModTime().After(dstInfo.ModTime()) {
            return copyFile(srcPath, dstPath)
        }

        return nil
    })
}

func copyFile(src, dst string) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    dstFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer dstFile.Close()

    _, err = io.Copy(dstFile, srcFile)
    if err != nil {
        return err
    }

    srcInfo, _ := os.Stat(src)
    return os.Chmod(dst, srcInfo.Mode())
}

func main() {
    myApp := app.NewWithID("folder-sync")
    myWin := myApp.NewWindow("🗂️ Folder Sync (Go + Fyne)")
    myWin.Resize(fyne.NewSize(500, 300))

    var srcPath, dstPath string

    srcEntry := widget.NewEntry()
    srcEntry.SetPlaceHolder("Source folder")

    dstEntry := widget.NewEntry()
    dstEntry.SetPlaceHolder("Destination folder")

    pickSrc := widget.NewButton("📁 Browse Source", func() {
        dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
            if uri != nil {
                srcPath = uri.Path()
                srcEntry.SetText(srcPath)
            }
        }, myWin)
    })

    pickDst := widget.NewButton("📁 Browse Destination", func() {
        dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
            if uri != nil {
                dstPath = uri.Path()
                dstEntry.SetText(dstPath)
            }
        }, myWin)
    })

    runSync := widget.NewButton("🔁 Start Sync", func() {
        if srcPath == "" || dstPath == "" {
            dialog.ShowError(fmt.Errorf("Select both source and destination"), myWin)
            return
        }

        err := syncFolders(srcPath, dstPath)
        if err != nil {
            dialog.ShowError(err, myWin)
        } else {
            dialog.ShowInformation("Success", "Folders synced successfully!", myWin)
        }
    })

    content := container.NewVBox(
        widget.NewLabel("🧊 Select folders to sync:"),
        srcEntry, pickSrc,
        dstEntry, pickDst,
        runSync,
    )

    myWin.SetContent(content)
    myWin.ShowAndRun()
}
