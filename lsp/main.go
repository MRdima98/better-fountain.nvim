package main

import (
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

func main() {
	commonlog.Configure(2, &fileLog)

	handler = protocol.Handler{
		Initialize:             initialize,
		Shutdown:               shutdown,
		TextDocumentColor:      textDocumentColor,
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

func textDocumentColor(context *glsp.Context, params *protocol.DocumentColorParams) ([]protocol.ColorInformation, error) {
	colors := make([]protocol.ColorInformation, 1)
	color := protocol.ColorInformation{
		Range: protocol.Range{
			Start: protocol.Position{Line: 2, Character: 3},
			End:   protocol.Position{Line: 2, Character: 10},
		},
		Color: protocol.Color{Red: 1.0, Green: 0.0, Blue: 0.0, Alpha: 0.0},
	}
	colors[0] = color
	return colors, nil
}

func textDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	hello := "wakanda"

    return []protocol.CompletionItem{
        {
            Label: "wakanda",
            InsertText: &hello,
            Detail: &hello,
        },
    }, nil
}
