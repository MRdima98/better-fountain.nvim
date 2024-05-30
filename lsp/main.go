package main

import (
	// "regexp"
	"strconv"

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

var characters []protocol.CompletionItem

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

	// change := protocol.TextDocumentSyncKindFull
	capabilities.CompletionProvider = &protocol.CompletionOptions{}
	// capabilities.TextDocumentSync = &protocol.TextDocumentSyncOptions{
	// 	Change: &change,
	// }

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

func didOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	// pattern := `\n\n\b[A-Z]+\b`
	// regex := regexp.MustCompile(pattern)
	// matches := regex.FindAllString(params.TextDocument.Text, -1)
	// var uniqueMatches []string
	// for _, match := range matches {
	// 	isPresent := false
	// 	for _, uniqueMatch := range uniqueMatches {
	// 		if uniqueMatch == match {
	// 			isPresent = true
	// 		}
	// 	}
	// 	if !isPresent {
	// 		uniqueMatches = append(uniqueMatches, match)
	// 	}
	// }

	// for _, match := range uniqueMatches {
	// 	tmp := match[2:]
	// 	tmp += "\n"
	// 	characters = append(characters, protocol.CompletionItem{
	// 		Label:      tmp,
	// 		InsertText: &tmp,
	// 	})
	// }
	// characters = UpdateCompletionList(params.TextDocument.Text)

	return nil
}

func didChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	commonlog.NewInfoMessage(0, "Log outside loop").Send()
	_, ok := params.ContentChanges[0].(protocol.TextDocumentContentChangeEventWhole)
	commonlog.NewInfoMessage(0, "Whole assertion: "+strconv.FormatBool(ok)).Send()

	for _, el := range params.ContentChanges {
		tmp, _ := el.(protocol.TextDocumentContentChangeEvent)
		commonlog.NewInfoMessage(0, tmp.Text).Send()
	}

	return nil
}

func didClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	return nil
}

func didSave(context *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
	if params.Text != nil {
		commonlog.NewInfoMessage(0, "Maybe: "+*params.Text).Send()
	}
	return nil
}
