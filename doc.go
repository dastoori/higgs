// Package higgs is a tiny cross-platform Go library to hide or unhide files
// and directories.
//
// Supported OSs: all unix based OSs (tested on Ubuntu, macOS) and Windows.
//
// How hiding works differs per platform:
//
//   - Unix (Linux, *BSD, macOS): a file is hidden by renaming it with a
//     leading dot (foo.txt -> .foo.txt). Because the file is renamed,
//     FileHide.Path (and the string returned by the Hide/Unhide helpers) is
//     updated to the new name.
//   - Windows: a file is hidden by toggling the FILE_ATTRIBUTE_HIDDEN
//     attribute; it is not renamed, so FileHide.Path is unchanged.
//   - macOS native (opt-in via DarwinNativeOption): higgs sets/clears the
//     UF_HIDDEN file flag instead of renaming, so FileHide.Path is unchanged
//     (like Windows). This option is only supported on macOS.
//
// Hide, Unhide and the FileHide methods are idempotent: hiding an
// already-hidden file (or unhiding an already-visible one) is a no-op.
package higgs
