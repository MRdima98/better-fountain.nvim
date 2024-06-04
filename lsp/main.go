package main

import (
	"github.com/tliron/commonlog"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"

	_ "github.com/tliron/commonlog/simple"
)

const lsName = "better-fountain"

var fileLog = "/tmp/lsp.log"

var version = "0.0.1"
var handler protocol.Handler

var complitionList []protocol.CompletionItem

func main() {
	commonlog.Configure(2, &fileLog)

	handler = protocol.Handler{
		Initialize:             initialize,
		Shutdown:               shutdown,
		TextDocumentDidOpen:    didOpen,
		TextDocumentDidChange:  didChange,
		TextDocumentDidClose:   didClose,
		TextDocumentDidSave:    didSave,
		TextDocumentCompletion: textDocumentCompletion,
	}

	server := server.NewServer(&handler, lsName, true)

	server.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	commonlog.NewInfoMessage(0, "Initializing server...").Send()

	capabilities := handler.CreateServerCapabilities()

	change := protocol.TextDocumentSyncKindFull
	capabilities.CompletionProvider = &protocol.CompletionOptions{}
	capabilities.TextDocumentSync = &protocol.TextDocumentSyncOptions{
		Change: &change,
	}

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
	return complitionList, nil
}

func didOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	complitionList = UpdateCompletionList(params.TextDocument.Text)
	return nil
}

func didChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	content, _ := params.ContentChanges[0].(protocol.TextDocumentContentChangeEventWhole)
	complitionList = UpdateCompletionList(content.Text)

	return nil
}

func didClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	return nil
}

func didSave(context *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
	return nil
}
