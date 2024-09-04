# Better-fountain port 
Minimal port of better-fountain to Neovim, written in Golang.

## Why?
I like to text edit in neovim, I like to write screenplays and I like a neat LSP. 

## Example 
![Alt Text](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExZjJqMDl6dHE3M2tnemZjZXpuNmJjazc4ZWFsamVocmExdTd6Y3hjbiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/yPfFGAeNPtxwZFyY7C/giphy.gif)

## Install
Using lazy:
```
{
'MRdima98/better-fountain.nvim',
  dependencies = { 'kblin/vim-fountain' }
}
```
vim-fountain provides syntax highlight, while my plugin provides suggestions for characters and scene headers.

## Start
This snippet will set up your lsp to start when you enter a fountain file.
```
vim.api.nvim_create_autocmd({"BufEnter", "BufWinEnter"}, {
    pattern = { "*.fountain" },
    callback = function ()
        vim.lsp.start({
            name = "better-fountain",
            cmd = { "path/to/executable" },
        })
    end
})
```
Your executable should be in the folder .local/share/nvim.

## Small catch
It won't work for Windows... even if you compile the file with go run... but I won't support that so just use [vscode](https://marketplace.visualstudio.com/items?itemName=piersdeseilligny.betterfountain).

## Where's the rest? 
Golang regexp is limited, until it gets better or I loose too much time writing a regex library or a parser, I won't implement functionalities I don't really need.
