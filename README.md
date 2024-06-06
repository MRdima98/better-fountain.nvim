# Better-fountain port 
Minimal port of better-fountain to Neovim, writting in Golang.

# Why?
I like to text edit in neovim, I like to write screenplays and I like a neat LSP. 

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
vim.api.nvim_create_autocmd({"BufEnter", "BufWinEnter"}, {
    pattern = { "*.fountain" },
    callback = function ()
        vim.lsp.start({
            name = "better-fountain",
            cmd = { "path/to/executable" },
        })
    end
})
Your executable should be in the folder nvim-:

# Where's the rest? 
Golang regexp is limited, until it gets better or I loose too much time writing a regex library, I won't implement functionalities I don't really need.
