import { _, Features, _Connection } from './server';
import { SemanticTokensBuilder } from './semanticTokens';
import type { WorkDoneProgressReporter, WorkDoneProgressServerReporter, ResultProgressReporter } from './progress';
export * from 'vscode-languageserver-protocol/';
export { WorkDoneProgressReporter, WorkDoneProgressServerReporter, ResultProgressReporter };
export { SemanticTokensBuilder };
import { TextDocuments, TextDocumentsConfiguration, TextDocumentChangeEvent, TextDocumentWillSaveEvent } from './textDocuments';
export { TextDocuments, TextDocumentsConfiguration, TextDocumentChangeEvent, TextDocumentWillSaveEvent };
export * from './server';
import { DiagnosticsFeatureShape } from './proposed.diagnostic';
import { TypeHierarchyFeatureShape } from './proposed.typeHierarchy';
import { InlineValuesFeatureShape } from './proposed.inlineValues';
import { NotebooksFeatureShape, NotebookDocuments as _NotebookDocuments } from './proposed.notebooks';
export declare namespace ProposedFeatures {
    const all: Features<_, _, _, _, _, _, DiagnosticsFeatureShape & TypeHierarchyFeatureShape & InlineValuesFeatureShape, NotebooksFeatureShape>;
    type Connection = _Connection<_, _, _, _, _, _, DiagnosticsFeatureShape & TypeHierarchyFeatureShape & InlineValuesFeatureShape, NotebooksFeatureShape>;
    const NotebookDocuments: typeof _NotebookDocuments;
}
