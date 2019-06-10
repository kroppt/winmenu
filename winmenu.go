package winmenu

import (
	"syscall"
	"unsafe"
)

var (
	moduser32          = syscall.NewLazyDLL("user32.dll")
	procCreateMenu     = moduser32.NewProc("CreateMenu")
	procInsertMenuItem = moduser32.NewProc("InsertMenuItemW")
)

// HMenu is a handle to a menu.
// (https://docs.microsoft.com/en-us/windows/desktop/WinProg/windows-data-types#HMENU)
type HMenu uintptr

// HBitmap is a handle to a bitmap.
// (https://docs.microsoft.com/en-us/windows/desktop/WinProg/windows-data-types#HBITMAP)
type HBitmap uintptr

// MaskFlag is a MenuItemInfo flag.
type MaskFlag uint32

// MenuItemInfo flags. Indicates the members to be retrieved or set. This
// member can be one or more of the following values.
const (
	// Retrieves or sets the hbmpItem member.
	MIIM_BITMAP MaskFlag = 0x00000080
	// Retrieves or sets the hbmpChecked and hbmpUnchecked members.
	MIIM_CHECKMARKS MaskFlag = 0x00000008
	// Retrieves or sets the dwItemData member.
	MIIM_DATA MaskFlag = 0x00000020
	// Retrieves or sets the fType member.
	MIIM_FTYPE MaskFlag = 0x00000100
	// Retrieves or sets the wID member.
	MIIM_ID MaskFlag = 0x00000002
	// Retrieves or sets the fState member.
	MIIM_STATE MaskFlag = 0x00000001
	// Retrieves or sets the dwTypeData member.
	MIIM_STRING MaskFlag = 0x00000040
	// Retrieves or sets the hSubMenu member.
	MIIM_SUBMENU MaskFlag = 0x00000004
	// Retrieves or sets the fType and dwTypeData members.
	// MIIM_TYPE is replaced by MIIM_BITMAP, MIIM_FTYPE, and MIIM_STRING.
	MIIM_TYPE MaskFlag = 0x00000010
)

// TypeFlag is a fType flag.
type TypeFlag uint32

// The menu item type. This member can be one or more of the following values.
const (
	// Displays the menu item using a bitmap. The low-order word of the
	// dwTypeData member is the bitmap handle, and the cch member is ignored.
	// MFT_BITMAP is replaced by MIIM_BITMAP and hbmpItem.
	MFT_BITMAP TypeFlag = 0x00000004
	// Places the menu item on a new line (for a menu bar) or in a new column
	// (for a drop-down menu, submenu, or shortcut menu). For a drop-down menu,
	// submenu, or shortcut menu, a vertical line separates the new column from
	// the old.
	MFT_MENUBARBREAK TypeFlag = 0x00000020
	// Places the menu item on a new line (for a menu bar) or in a new column
	// (for a drop-down menu, submenu, or shortcut menu). For a drop-down menu,
	// submenu, or shortcut menu, the columns are not separated by a vertical
	// line.
	MFT_MENUBREAK TypeFlag = 0x00000040
	// Assigns responsibility for drawing the menu item to the window that owns
	// the menu. The window receives a WM_MEASUREITEM
	// (https://msdn.microsoft.com/en-us/library/Bb775925(v=VS.85).aspx)
	// message before the menu is displayed for the first time, and a
	// WM_DRAWITEM
	// (https://msdn.microsoft.com/en-us/library/Bb775923(v=VS.85).aspx)
	// message whenever the appearance of the menu item must be updated. If
	// this value is specified, the dwTypeData member contains an application-
	// defined value.
	MFT_OWNERDRAW TypeFlag = 0x00000100
	// Displays selected menu items using a radio-button mark instead of a check
	// mark if the hbmpChecked member is nil.
	MFT_RADIOCHECK TypeFlag = 0x00000200
	// Specifies that menus cascade right-to-left (the default is left-to-
	// right). This is used to support right-to-left languages, such as Arabic
	// and Hebrew.
	MFT_RIGHTORDER TypeFlag = 0x00002000
	// Specifies that the menu item is a separator. A menu item separator
	// appears as a horizontal dividing line. The dwTypeData and cch members are
	// ignored. This value is valid only in a drop-down menu, submenu, or
	// shortcut menu.
	MFT_SEPARATOR TypeFlag = 0x00000800
	// Displays the menu item using a text string. The dwTypeData member is the
	// pointer to a null-terminated string, and the cch member is the length of
	// the string.
	// MFT_STRING is replaced by MIIM_STRING.
	MFT_STRING TypeFlag = 0x00000000
)

// StateFlag is a fState flag.
type StateFlag uint32

// The menu item state. This member can be one or more of these values. Set
// fMask to MIIM_STATE to use fState.
const (
	// Checks the menu item. For more information about selected menu items, see
	// the hbmpChecked member.
	MFS_CHECKED StateFlag = 0x00000008
	// Specifies that the menu item is the default. A menu can contain only one
	// default menu item, which is displayed in bold.
	MFS_DEFAULT StateFlag = 0x00001000
	// Disables the menu item and grays it so that it cannot be selected. This
	// is equivalent to MFS_GRAYED.
	MFS_DISABLED StateFlag = 0x00000003
	// Enables the menu item so that it can be selected. This is the default
	// state.
	MFS_ENABLED StateFlag = 0x00000000
	// Disables the menu item and grays it so that it cannot be selected. This
	// is equivalent to MFS_DISABLED.
	MFS_GRAYED StateFlag = 0x00000003
	// Highlights the menu item.
	MFS_HILITE StateFlag = 0x00000080
	// Unchecks the menu item. For more information about clear menu items, see
	// the hbmpChecked member.
	MFS_UNCHECKED StateFlag = 0x00000000
	// Removes the highlight from the menu item. This is the default state.
	MFS_UNHILITE StateFlag = 0x00000000
)

// A handle to the bitmap to be displayed, or it can be one of the values in the
// following table. It is used when the MIIM_BITMAP flag is set in the fMask
// member.
const (
	// A bitmap that is drawn by the window that owns the menu. The application
	// must process the WM_MEASUREITEM
	// (https://msdn.microsoft.com/en-us/library/Bb775925(v=VS.85).aspx)
	// and WM_DRAWITEM
	// (https://msdn.microsoft.com/en-us/library/Bb775923(v=VS.85).aspx)
	// messages.
	// TODO figure out how to cast -1 as a pointer
	// HBMMENU_CALLBACK uintptr = -1
	// Close button for the menu bar.
	HBMMENU_MBAR_CLOSE HBitmap = 5
	// Disabled close button for the menu bar.
	HBMMENU_MBAR_CLOSE_D HBitmap = 6
	// Minimize button for the menu bar.
	HBMMENU_MBAR_MINIMIZE HBitmap = 3
	// Disabled minimize button for the menu bar.
	HBMMENU_MBAR_MINIMIZE_D HBitmap = 7
	// Restore button for the menu bar.
	HBMMENU_MBAR_RESTORE HBitmap = 2
	// Close button for the submenu.
	HBMMENU_POPUP_CLOSE HBitmap = 8
	// Maximize button for the submenu.
	HBMMENU_POPUP_MAXIMIZE HBitmap = 10
	// Minimize button for the submenu.
	HBMMENU_POPUP_MINIMIZE HBitmap = 11
	// Restore button for the submenu.
	HBMMENU_POPUP_RESTORE HBitmap = 9
	// Windows icon or the icon of the window specified in dwItemData.
	HBMMENU_SYSTEM HBitmap = 1
)

// MenuItemInfo contains information about a menu item.
//
// Remarks:
//   - The MENUITEMINFO structure is used with the GetMenuItemInfo
// (https://msdn.microsoft.com/en-us/library/ms647980(v=VS.85).aspx),
// InsertMenuItem
// (https://msdn.microsoft.com/en-us/library/ms647988(v=VS.85).aspx),
// and SetMenuItemInfo
// (https://msdn.microsoft.com/en-us/library/ms648001(v=VS.85).aspx) functions.
//   - The menu can display items using text, bitmaps, or both.
// Requirements:
//   - Minimum supported client = Windows 2000 Professional [desktop apps only]
//   - Minimum supported server = Windows 2000 Server [desktop apps only]
//   - Header = winuser.h (include Windows.h)
// (https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-menuiteminfow)
type MenuItemInfo struct {
	// The size of the structure, in bytes.
	cbSize uint32 // set by caller
	// Indicates the members to be retrieved or set.
	fMask MaskFlag
	// The menu item type.
	// The MFT_BITMAP, MFT_SEPARATOR, and MFT_STRING values cannot be combined
	// with one another. Set fMask to MIIM_TYPE to use fType.
	// fType is used only if fMask has a value of MIIM_FTYPE.
	fType TypeFlag
	// The menu item state.
	// Set fMask to MIIM_STATE to use fState.
	fState StateFlag
	// An application-defined value that identifies the menu item.
	// Set fMask to MIIM_ID to use wID.
	wID uint32
	// A handle to the drop-down menu or submenu associated with the menu item.
	// If the menu item is not an item that opens a drop-down menu or submenu,
	// this member is NULL. Set fMask to MIIM_SUBMENU to use hSubMenu.
	hSubMenu HMenu
	// A handle to the bitmap to display next to the item if it is selected.
	// If this member is NULL, a default bitmap is used. If the MFT_RADIOCHECK
	// type value is specified, the default bitmap is a bullet. Otherwise, it is
	// a check mark. Set fMask to MIIM_CHECKMARKS to use hbmpChecked.
	hbmpChecked HBitmap
	// A handle to the bitmap to display next to the item if it is not selected.
	// If this member is NULL, no bitmap is used. Set fMask to MIIM_CHECKMARKS
	// to use hbmpUnchecked.
	hbmpUnchecked HBitmap
	// An application-defined value associated with the menu item. Set fMask to
	// MIIM_DATA to use dwItemData.
	dwItemData *uint64
	// The contents of the menu item. The meaning of this member depends on the
	// value of fType and is used only if the MIIM_TYPE flag is set in the
	// fMask member.
	dwTypeData *uint16
	// The length of the menu item text, in characters, when information is
	// received about a menu item of the MFT_STRING type. However, cch is used
	// only if the MIIM_TYPE flag is set in the fMask member and is zero
	// otherwise. Also, cch is ignored when the content of a menu item is set by
	// calling SetMenuItemInfo
	// (https://msdn.microsoft.com/en-us/library/ms648001(v=VS.85).aspx).
	//
	// Note that, before calling GetMenuItemInfo
	// (https://msdn.microsoft.com/en-us/library/ms647980(v=VS.85).aspx),
	// the application must set cch to the length of the buffer pointed to by
	// the dwTypeData member. If the retrieved menu item is of type MFT_STRING
	// (as indicated by the fType member), then GetMenuItemInfo changes cch to
	// the length of the menu item text. If the retrieved menu item is of some
	// other type, GetMenuItemInfo sets the cch field to zero.
	//
	// The cch member is used when the MIIM_STRING flag is set in the fMask
	// member.
	cch uint32
	// A handle to the bitmap to be displayed, or it can be one of the defined
	// HBitmap constants. It is used when the MIIM_BITMAP flag is set in the
	// fMask member.
	hbmpItem HBitmap
}

// CreateMenu creates a menu.
// (https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-createmenu)
func CreateMenu() (hMenu HMenu, ok bool) {
	ret, _, _ := procCreateMenu.Call()
	return HMenu(ret), ret != 0
}

// InsertMenuItem inserts a new menu item at the specified position in a menu.
// (https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-insertmenuitemw)
func (hMenu HMenu) InsertMenuItem(item uint32, fByPosition bool, lpmi *MenuItemInfo) (ok bool) {
	byPos := 0
	if fByPosition {
		byPos = 1
	}
	lpmi.cbSize = uint32(unsafe.Sizeof(*lpmi))
	ret, _, _ := procInsertMenuItem.Call(uintptr(hMenu), uintptr(item), uintptr(byPos), uintptr(unsafe.Pointer(lpmi)))
	return ret != 0
}
