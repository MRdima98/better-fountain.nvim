# Better-fountain port 
Minimal port of better-fountain to Neovim, writting in Golang.

# Why?
I like to text edit in neovim, I like to write screenplays and I like a neat LSP. 

# Example 
![Alt Text](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExZjJqMDl6dHE3M2tnemZjZXpuNmJjazc4ZWFsamVocmExdTd6Y3hjbiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/yPfFGAeNPtxwZFyY7C/giphy.gif)

# Install
Using lazy:
```
-- init.lua:
    {
    'MRdima98/better-fountain.nvim',
      dependencies = { 'kblin/vim-fountain' }
    }
```
vim-fountain provides syntax highlight, while my plugin provides suggestions for characters and scene headers.

# Start
This will set up your lsp to start when you ender a fountain file.
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

# Where's the rest? 
Golang regexp is limited, until it gets better or I loose too much time writing a regex library, I won't implement functionalities I don't really need.
