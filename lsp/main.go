package main

import (
	"regexp"
	"github.com/tliron/commonlog"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"

	_ "github.com/tliron/commonlog/simple"
)

const lsName = "gotem"

var fileLog = "/tmp/lsp.log"

var version string = "0.0.1"
var handler protocol.Handler

var characters []protocol.CompletionItem


func main() {
    commonlog.Configure(2, &fileLog)

    handler = protocol.Handler{
        Initialize:             initialize,
        Shutdown:               shutdown,
        TextDocumentDidOpen: didOpen,
        TextDocumentDidChange: didChange,
        TextDocumentDidClose: didClose,
        TextDocumentCompletion: textDocumentCompletion,
    }

    server := server.NewServer(&handler, lsName, true)

    server.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
    commonlog.NewInfoMessage(0, "Initializing server...")

    capabilities := handler.CreateServerCapabilities()

    capabilities.CompletionProvider = &protocol.CompletionOptions{}

    return protocol.InitializeResult{
        Capabilities: capabilities,
        ServerInfo: &protocol.InitializeResultServerInfo{
            Name:    lsName,
            Version: &version,
        },
    }, nil
}

func shutdown(context *glsp.Context) error {
    return nil
}

func textDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
    return characters, nil
}

func didOpen (context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	pattern := `\n\n\b[A-Z]+\b`
    regex := regexp.MustCompile(pattern)
    matches := regex.FindAllString(params.TextDocument.Text, -1)
    for _, match := range matches {
        tmp := match[2:]
        tmp += "\n"
        characters = append(characters, protocol.CompletionItem{
            Label: tmp,
            InsertText: &tmp,
        })
    }

    return nil
}

func didChange (context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
    hello := "CALEEN"
    present := false
    for _, char := range characters {
        if "CALEEN" == char.Label {
            present = true
        }
    }
    if !present {
        characters = append(characters, protocol.CompletionItem{
            Label: hello,
            InsertText: &hello,
        })
    }
    return nil
}

func didClose (context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
    return nil
}
