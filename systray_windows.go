/*
 * Copyright (C) 2019 The Systray Authors. All Rights Reserved.
 */

package systray

import (
	"errors"
	"path/filepath"
	"syscall"
	"unsafe"
)

//goland:noinspection GoSnakeCaseUsage,GoUnusedConst,SpellCheckingInspection
const (
	WM_LBUTTONUP     = 0x0202
	WM_LBUTTONDBLCLK = 0x0203
	WM_RBUTTONUP     = 0x0205
	WM_USER          = 0x0400
	WM_TRAYICON      = WM_USER + 69

	WS_EX_APPWINDOW     = 0x00040000
	WS_OVERLAPPEDWINDOW = 0x00000000 | 0x00C00000 | 0x00080000 | 0x00040000 | 0x00020000 | 0x00010000
	CW_USEDEFAULT       = 0x80000000

	NIM_ADD        = 0x00000000
	NIM_MODIFY     = 0x00000001
	NIM_DELETE     = 0x00000002
	NIM_SETVERSION = 0x00000004

	NIF_MESSAGE = 0x00000001
	NIF_ICON    = 0x00000002
	NIF_TIP     = 0x00000004
	NIF_STATE   = 0x00000008
	NIF_INFO    = 0x00000010

	NIS_HIDDEN = 0x00000001

	NIIF_NONE               = 0x00000000
	NIIF_INFO               = 0x00000001
	NIIF_WARNING            = 0x00000002
	NIIF_ERROR              = 0x00000003
	NIIF_USER               = 0x00000004
	NIIF_NOSOUND            = 0x00000010
	NIIF_LARGE_ICON         = 0x00000020
	NIIF_RESPECT_QUIET_TIME = 0x00000080
	NIIF_ICON_MASK          = 0x0000000F

	IMAGE_BITMAP    = 0
	IMAGE_ICON      = 1
	LR_LOADFROMFILE = 0x00000010
	LR_DEFAULTSIZE  = 0x00000040

	IDC_ARROW     = 32512
	COLOR_WINDOW  = 5
	COLOR_BTNFACE = 15

	GWLP_USERDATA       = -21
	WS_CLIPSIBLINGS     = 0x04000000
	WS_EX_CONTROLPARENT = 0x00010000

	HWND_MESSAGE       = ^HWND(2)
	NOTIFYICON_VERSION = 4

	IDI_APPLICATION = 32512
	WM_APP          = 32768
	WM_COMMAND      = 273

	MenuItemMsgID       = WM_APP + 1024
	NotifyIconMessageId = WM_APP + iota

	MF_STRING       = 0x00000000
	MF_ENABLED      = 0x00000000
	MF_GRAYED       = 0x00000001
	MF_DISABLED     = 0x00000002
	MF_SEPARATOR    = 0x00000800
	MF_CHECKED      = 0x00000008
	MF_MENUBARBREAK = 0x00000020

	TPM_LEFTALIGN = 0x0000
	WM_NULL       = 0
)

//goland:noinspection GoUnusedGlobalVariable
var (
	kernel32         = syscall.MustLoadDLL("kernel32")
	GetModuleHandle  = kernel32.MustFindProc("GetModuleHandleW")
	GetConsoleWindow = kernel32.MustFindProc("GetConsoleWindow")
	GetLastError     = kernel32.MustFindProc("GetLastError")

	shell32         = syscall.MustLoadDLL("shell32.dll")
	ShellNotifyIcon = shell32.MustFindProc("Shell_NotifyIconW")

	user32 = syscall.MustLoadDLL("user32.dll")

	GetMessage       = user32.MustFindProc("GetMessageW")
	IsDialogMessage  = user32.MustFindProc("IsDialogMessageW")
	TranslateMessage = user32.MustFindProc("TranslateMessage")
	DispatchMessage  = user32.MustFindProc("DispatchMessageW")

	ShowWindow       = user32.MustFindProc("ShowWindow")
	UpdateWindow     = user32.MustFindProc("UpdateWindow")
	DefWindowProc    = user32.MustFindProc("DefWindowProcW")
	RegisterClassEx  = user32.MustFindProc("RegisterClassExW")
	GetDesktopWindow = user32.MustFindProc("GetDesktopWindow")
	CreateWindowEx   = user32.MustFindProc("CreateWindowExW")

	CreatePopupMenu         = user32.MustFindProc("CreatePopupMenu")
	procAppendMenuW         = user32.MustFindProc("AppendMenuW")
	procGetCursorPos        = user32.MustFindProc("GetCursorPos")
	procSetForegroundWindow = user32.MustFindProc("SetForegroundWindow")
	procTrackPopupMenu      = user32.MustFindProc("TrackPopupMenu")
	procPostMessage         = user32.MustFindProc("PostMessageW")

	LoadImage  = user32.MustFindProc("LoadImageW")
	LoadIcon   = user32.MustFindProc("LoadIconW")
	LoadCursor = user32.MustFindProc("LoadCursorW")
)

//goland:noinspection SpellCheckingInspection
type NOTIFYICONDATA struct {
	CbSize           uint32
	HWnd             HWND
	UID              uint32
	UFlags           uint32
	UCallbackMessage uint32
	HIcon            HICON
	SzTip            [128]uint16
	DwState          uint32
	DwStateMask      uint32
	SzInfo           [256]uint16
	UVersion         uint32
	SzInfoTitle      [64]uint16
	DwInfoFlags      uint32
	GuidItem         GUID
	HBalloonIcon     HICON
}

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

//goland:noinspection SpellCheckingInspection
type WNDCLASSEX struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       HICON
}

type MSG struct {
	HWnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

type POINT struct {
	X, Y int32
}

//goland:noinspection SpellCheckingInspection
type (
	HANDLE    uintptr
	HINSTANCE HANDLE
	HCURSOR   HANDLE
	HICON     HANDLE
	HWND      HANDLE
	HGDIOBJ   HANDLE
	HBRUSH    HGDIOBJ
)

//goland:noinspection SpellCheckingInspection
type HMENU HANDLE

type WindowProc func(hWnd HWND, msg uint32, wParam, lParam uintptr) uintptr

type MenuItem struct {
	Label string

	Disabled  bool
	Checked   bool
	BarBreak  bool
	Separator bool

	OnClick func()
}

type Systray struct {
	id     uint32
	hWnd   HWND
	hInst  HINSTANCE
	lClick func()
	rClick func()

	Menu []*MenuItem
}

func New() (*Systray, error) {
	ni := &Systray{lClick: func() {}, rClick: func() {}}

	MainClassName := "MainForm"
	ni.hInst, _ = RegisterWindow(MainClassName, ni.WinProc)

	mainClass, _ := syscall.UTF16PtrFromString(MainClassName)
	mHWnd, _, _ := CreateWindowEx.Call(
		WS_EX_CONTROLPARENT,
		uintptr(unsafe.Pointer(mainClass)),
		0,
		WS_OVERLAPPEDWINDOW|WS_CLIPSIBLINGS,
		CW_USEDEFAULT,
		CW_USEDEFAULT,
		CW_USEDEFAULT,
		CW_USEDEFAULT,
		0,
		0,
		0,
		0)
	if mHWnd == 0 {
		return nil, errors.New("create main win failed")
	}

	NotifyIconClassName := "NotifyIconForm"
	_, _ = RegisterWindow(NotifyIconClassName, ni.WinProc)

	iconClass, _ := syscall.UTF16PtrFromString(NotifyIconClassName)
	hWnd, _, _ := CreateWindowEx.Call(
		0,
		uintptr(unsafe.Pointer(iconClass)),
		0,
		0,
		0,
		0,
		0,
		0,
		uintptr(HWND_MESSAGE),
		0,
		0,
		0)
	if hWnd == 0 {
		return nil, errors.New("create notify win failed")
	}

	ni.hWnd = HWND(hWnd) // Important to keep this inside struct.

	nid := NOTIFYICONDATA{
		HWnd:             HWND(hWnd),
		UFlags:           NIF_MESSAGE | NIF_STATE,
		DwState:          NIS_HIDDEN,
		DwStateMask:      NIS_HIDDEN,
		UCallbackMessage: NotifyIconMessageId,
	}
	nid.CbSize = uint32(unsafe.Sizeof(nid))

	ret, _, _ := ShellNotifyIcon.Call(NIM_ADD, uintptr(unsafe.Pointer(&nid)))
	if ret == 0 {
		return nil, errors.New("shell notify create failed")
	}

	nid.UVersion = NOTIFYICON_VERSION

	ret, _, _ = ShellNotifyIcon.Call(NIM_SETVERSION, uintptr(unsafe.Pointer(&nid)))
	if ret == 0 {
		return nil, errors.New("shell notify version failed")
	}

	ni.id = nid.UID
	return ni, nil
}

func (p *Systray) HWND() HWND {
	return p.hWnd
}

// AppendMenu add menu item.
func (p *Systray) AppendMenu(label string, onclick func()) {
	p.Menu = append(p.Menu, &MenuItem{Label: label, OnClick: onclick})
}

// AppendSeparator to the menu.
func (p *Systray) AppendSeparator() {
	p.Menu = append(p.Menu, &MenuItem{Separator: true})
}

func (p *Systray) Stop() error {
	nid := NOTIFYICONDATA{
		UID:  p.id,
		HWnd: p.hWnd,
	}
	nid.CbSize = uint32(unsafe.Sizeof(nid))

	ret, _, _ := ShellNotifyIcon.Call(NIM_DELETE, uintptr(unsafe.Pointer(&nid)))
	if ret == 0 {
		return errors.New("shell notify delete failed")
	}
	return nil
}

func MakeIntResource(id uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(id)))
}

// Show shows custom tray icon.
func (p *Systray) Show(iconResID uint16, hint string) error {
	icon, _, _ := LoadIcon.Call(uintptr(p.hInst), uintptr(unsafe.Pointer(MakeIntResource(iconResID))))
	if icon == 0 {
		icon, _, _ = LoadIcon.Call(0, uintptr(IDI_APPLICATION))
	}

	err := p.SetIcon(HICON(icon))
	if err != nil {
		return err
	}
	err = p.SetTooltip(hint)
	if err != nil {
		return err
	}
	return p.SetVisible(true)
}

func loadIconFile(file string) (HICON, error) {
	path, err := filepath.Abs(file)
	if err != nil {
		return 0, err
	}
	icon, err := NewIconFromFile(path)
	if err != nil {
		return 0, err
	}
	return HICON(icon), nil
}

// ShowCustom shows custom tray icon.
func (p *Systray) ShowCustom(file string, hint string) error {
	hIcon, err := loadIconFile(file)
	if err != nil {
		icon, _, _ := LoadIcon.Call(0, uintptr(IDI_APPLICATION))
		hIcon = HICON(icon)
	}

	err = p.SetIcon(hIcon)
	if err != nil {
		return err
	}
	err = p.SetTooltip(hint)
	if err != nil {
		return err
	}
	return p.SetVisible(true)
}

func (p *Systray) OnClick(fn func()) {
	p.lClick = fn
}

func (p *Systray) OnRightClick(fn func()) {
	p.rClick = fn
}

func (p *Systray) SetTooltip(tooltip string) error {
	nid := NOTIFYICONDATA{
		UID:  p.id,
		HWnd: p.hWnd,
	}
	nid.CbSize = uint32(unsafe.Sizeof(nid))

	nid.UFlags = NIF_TIP
	tip, _ := syscall.UTF16FromString(tooltip)
	copy(nid.SzTip[:], tip)

	ret, _, _ := ShellNotifyIcon.Call(NIM_MODIFY, uintptr(unsafe.Pointer(&nid)))
	if ret == 0 {
		return errors.New("shell notify tooltip failed")
	}
	return nil
}

func (p *Systray) ShowMessage(title, msg string, bigIcon bool) error {
	nid := NOTIFYICONDATA{
		UID:  p.id,
		HWnd: p.hWnd,
	}
	if bigIcon {
		nid.DwInfoFlags = NIIF_USER
	}

	nid.CbSize = uint32(unsafe.Sizeof(nid))

	nid.UFlags = NIF_INFO
	t, _ := syscall.UTF16FromString(title)
	m, _ := syscall.UTF16FromString(msg)
	copy(nid.SzInfoTitle[:], t)
	copy(nid.SzInfo[:], m)

	ret, _, _ := ShellNotifyIcon.Call(NIM_MODIFY, uintptr(unsafe.Pointer(&nid)))
	if ret == 0 {
		return errors.New("shell notify tooltip failed")
	}
	return nil
}

func (p *Systray) SetVisible(visible bool) error {
	nid := NOTIFYICONDATA{
		UID:  p.id,
		HWnd: p.hWnd,
	}
	nid.CbSize = uint32(unsafe.Sizeof(nid))

	nid.UFlags = NIF_STATE
	nid.DwStateMask = NIS_HIDDEN
	if !visible {
		nid.DwState = NIS_HIDDEN
	}

	ret, _, _ := ShellNotifyIcon.Call(NIM_MODIFY, uintptr(unsafe.Pointer(&nid)))
	if ret == 0 {
		return errors.New("shell notify tooltip failed")
	}
	return nil
}

func (p *Systray) SetIcon(hIcon HICON) error {
	nid := NOTIFYICONDATA{
		UID:  p.id,
		HWnd: p.hWnd,
	}
	nid.CbSize = uint32(unsafe.Sizeof(nid))

	nid.UFlags = NIF_ICON
	if hIcon == 0 {
		nid.HIcon = 0
	} else {
		nid.HIcon = hIcon
	}

	ret, _, _ := ShellNotifyIcon.Call(NIM_MODIFY, uintptr(unsafe.Pointer(&nid)))
	if ret == 0 {
		return errors.New("shell notify icon failed")
	}
	return nil
}

func (p *Systray) WinProc(hWnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case NotifyIconMessageId:
		if lParam == WM_LBUTTONUP {
			p.lClick()
			if len(p.Menu) > 0 {
				_ = p.displayMenu(p.Menu)
			}
		} else if lParam == WM_RBUTTONUP {
			p.rClick()
			if len(p.Menu) > 0 {
				_ = p.displayMenu(p.Menu)
			}
		}

	case WM_COMMAND:
		cmdMsgID := int(wParam & 0xffff)
		switch cmdMsgID {
		default:
			if cmdMsgID >= MenuItemMsgID && cmdMsgID < (MenuItemMsgID+len(p.Menu)) {
				itemIndex := cmdMsgID - MenuItemMsgID
				menuItem := p.Menu[itemIndex]
				menuItem.OnClick()
			}
		}
	}

	result, _, _ := DefWindowProc.Call(uintptr(hWnd), uintptr(msg), wParam, lParam)
	return result
}

func (p *Systray) Run() error {
	hWnd := p.hWnd
	var msg MSG
	for {
		rt, _, _ := GetMessage.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
		switch int(rt) {
		case 0:
			return nil
		case -1:
			return errors.New("run failed")
		}

		is, _, _ := IsDialogMessage.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&msg)))
		if is == 0 {
			_, _, _ = TranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
			_, _, _ = DispatchMessage.Call(uintptr(unsafe.Pointer(&msg)))
		}
	}
}

func (p *Systray) displayMenu(menuItems []*MenuItem) error {
	ret, _, _ := CreatePopupMenu.Call(0, 0, 0, 0)
	if ret == 0 {
		return errors.New("can not create menu")
	}
	menu := HMENU(ret)

	for index, item := range menuItems {
		var ret bool
		itemID := MenuItemMsgID + index
		flags := MF_STRING
		if item.Disabled {
			flags = flags | MF_GRAYED
		}
		if item.Checked {
			flags = flags | MF_CHECKED
		}
		if item.BarBreak {
			flags = flags | MF_MENUBARBREAK
		}
		if item.Separator {
			flags = flags | MF_SEPARATOR
		}

		ret = appendMenu(menu, uintptr(flags), uintptr(itemID), item.Label)
		if !ret {
			return errors.New("AppendMenu failed")
		}
	}

	x, y, ok := getCursorPos()
	if !ok {
		return errors.New("GetCursorPos failed")
	}

	if !setForegroundWindow(p.hWnd) {
		return errors.New("SetForegroundWindow failed")
	}

	if !trackPopupMenu(menu, TPM_LEFTALIGN, x, y-5, p.hWnd) {
		return errors.New("TrackPopupMenu failed")
	}

	if ret, _, _ := procPostMessage.Call(uintptr(p.hWnd), uintptr(WM_NULL), 0, 0); ret == 0 {
		return errors.New("PostMessage failed")
	}
	return nil
}

func trackPopupMenu(menu HMENU, flags uint, x, y int, wnd HWND) bool {
	ret, _, _ := procTrackPopupMenu.Call(
		uintptr(menu),
		uintptr(flags),
		uintptr(x),
		uintptr(y),
		0,
		uintptr(wnd),
		0,
	)
	return ret != 0
}

func setForegroundWindow(wnd HWND) bool {
	ret, _, _ := procSetForegroundWindow.Call(
		uintptr(wnd),
	)
	return ret != 0
}

func getCursorPos() (x, y int, ok bool) {
	pt := POINT{}
	ret, _, _ := procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	return int(pt.X), int(pt.Y), ret != 0
}

func appendMenu(menu HMENU, flags uintptr, id uintptr, text string) bool {
	t, _ := syscall.UTF16PtrFromString(text)
	ret, _, _ := procAppendMenuW.Call(
		uintptr(menu),
		flags,
		id,
		uintptr(unsafe.Pointer(t)),
	)
	return ret != 0
}

func NewIconFromFile(filePath string) (uintptr, error) {
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return 0, err
	}
	p, _ := syscall.UTF16PtrFromString(absFilePath)
	hIcon, _, _ := LoadImage.Call(
		0,
		uintptr(unsafe.Pointer(p)),
		IMAGE_ICON,
		0,
		0,
		LR_DEFAULTSIZE|LR_LOADFROMFILE)
	if hIcon == 0 {
		return 0, errors.New("load image failed: " + filePath)
	}
	return hIcon, nil
}

func RegisterWindow(name string, proc WindowProc) (HINSTANCE, error) {
	hInst, _, _ := GetModuleHandle.Call(0)
	if hInst == 0 {
		return 0, errors.New("get module handle failed")
	}
	hIcon, _, _ := LoadIcon.Call(0, uintptr(IDI_APPLICATION))
	if hIcon == 0 {
		return 0, errors.New("load icon failed")
	}
	hCursor, _, _ := LoadCursor.Call(0, uintptr(IDC_ARROW))
	if hCursor == 0 {
		return 0, errors.New("load cursor failed")
	}

	hi := HINSTANCE(hInst)

	var wc WNDCLASSEX
	wc.CbSize = uint32(unsafe.Sizeof(wc))
	wc.LpfnWndProc = syscall.NewCallback(proc)
	wc.HInstance = hi
	wc.HIcon = HICON(hIcon)
	wc.HCursor = HCURSOR(hCursor)
	wc.HbrBackground = COLOR_BTNFACE + 1
	wc.LpszClassName, _ = syscall.UTF16PtrFromString(name)

	atom, _, _ := RegisterClassEx.Call(uintptr(unsafe.Pointer(&wc)))
	if atom == 0 {
		return 0, errors.New("register class failed")
	}
	return hi, nil
}
