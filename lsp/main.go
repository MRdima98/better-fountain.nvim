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
        LogTrace: logTrace,
        TextDocumentSemanticTokensFull: tokens,
        TextDocumentDidOpen: didOpen,
        TextDocumentDidChange: didChange,
        TextDocumentDidClose: didClose,
        TextDocumentColor: color,
        TextDocumentDocumentHighlight: highlights,
        TextDocumentDocumentSymbol: documentSymbol,
		TextDocumentCompletion: textDocumentCompletion,
	}

	server := server.NewServer(&handler, lsName, true)

	server.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	commonlog.NewInfoMessage(0, "Initializing server...")

	capabilities := handler.CreateServerCapabilities()


    capabilities.CompletionProvider = &protocol.CompletionOptions{}
    capabilities.DocumentHighlightProvider = &protocol.DocumentHighlightOptions{ }

    capabilities.DocumentSymbolProvider = &protocol.DocumentSymbolOptions{ }

    capabilities.SemanticTokensProvider = &protocol.SemanticTokensOptions{
        Legend: protocol.SemanticTokensLegend{
            TokenTypes: []string{"type"},

            TokenModifiers: []string{"private"},
        },
        Full: true,
    }

    capabilities.ColorProvider = &protocol.DocumentColorOptions{}


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
	hello := "wakanda"

    return []protocol.CompletionItem{
        {
            Label: "wakanda",
            InsertText: &hello,
            Detail: &hello,
        },
    }, nil
}

func documentSymbol(context *glsp.Context, params *protocol.DocumentSymbolParams) (any, error) {
	commonlog.NewInfoMessage(0, "Symbols")
    // var symbols []protocol.DocumentSymbol

    return []protocol.DocumentSymbol{
        {
            Name: "symbols",
            Range: protocol.Range{
                Start: protocol.Position{ 
                    Line: 0,
                    Character: 0,
                },
                End: protocol.Position{ 
                    Line: 0,
                    Character: 5,
                },
            },
            Kind: 6,
        },
    }, nil
}

func highlights (context *glsp.Context, params *protocol.DocumentHighlightParams) ([]protocol.DocumentHighlight, error) {
    var hi []protocol.DocumentHighlight
    hello := protocol.DocumentHighlightKindWrite
    hi = append(hi, protocol.DocumentHighlight{
		Range: protocol.Range{
			Start: protocol.Position{Line: 0, Character: 0},
			End:   protocol.Position{Line: 0, Character: 10},
		},
        Kind: &hello,

    })
    return hi, nil
}


func didOpen (context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
    return nil
}

func didChange (context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
    return nil
}
func didClose (context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {

    return nil
}

func color (context *glsp.Context, params *protocol.DocumentColorParams) ([]protocol.ColorInformation, error) {
    return []protocol.ColorInformation{
        {
            Range: protocol.Range{
                Start: protocol.Position{Line: 0, Character: 0},
                End:   protocol.Position{Line: 0, Character: 10},
            },
            Color: protocol.Color{
                Red: 1.0,
                Green: 0,
                Blue: 0,
                Alpha: 0,
            },
        },
    }, nil
}

func logTrace (context *glsp.Context, params *protocol.LogTraceParams) error {
	commonlog.NewInfoMessage(4, *params.Verbose)
    return nil
}

func tokens (context *glsp.Context, params *protocol.SemanticTokensParams) (*protocol.SemanticTokens, error) {
    tokens := [] uint32{0,0,4,0,0}
    return &protocol.SemanticTokens{
        Data: tokens,
    },nil
}
