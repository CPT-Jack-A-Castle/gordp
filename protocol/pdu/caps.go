package pdu

import (
	"../../core"
	"log"
	"errors"
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240486.aspx
*/
type CapsType uint16
 
const /*CapsType*/ (
CAPSTYPE_GENERAL CapsType = 0x0001
CAPSTYPE_BITMAP = 0x0002
CAPSTYPE_ORDER = 0x0003
CAPSTYPE_BITMAPCACHE = 0x0004
CAPSTYPE_CONTROL = 0x0005
CAPSTYPE_ACTIVATION = 0x0007
CAPSTYPE_POINTER = 0x0008
CAPSTYPE_SHARE = 0x0009
CAPSTYPE_COLORCACHE = 0x000A
CAPSTYPE_SOUND = 0x000C
CAPSTYPE_INPUT = 0x000D
CAPSTYPE_FONT = 0x000E
CAPSTYPE_BRUSH = 0x000F
CAPSTYPE_GLYPHCACHE = 0x0010
CAPSTYPE_OFFSCREENCACHE = 0x0011
CAPSTYPE_BITMAPCACHE_HOSTSUPPORT = 0x0012
CAPSTYPE_BITMAPCACHE_REV2 = 0x0013
CAPSTYPE_VIRTUALCHANNEL = 0x0014
CAPSTYPE_DRAWNINEGRIDCACHE = 0x0015
CAPSTYPE_DRAWGDIPLUS = 0x0016
CAPSTYPE_RAIL = 0x0017
CAPSTYPE_WINDOW = 0x0018
CAPSETTYPE_COMPDESK = 0x0019
CAPSETTYPE_MULTIFRAGMENTUPDATE = 0x001A
CAPSETTYPE_LARGE_POINTER = 0x001B
CAPSETTYPE_SURFACE_COMMANDS = 0x001C
CAPSETTYPE_BITMAP_CODECS = 0x001D
CAPSSETTYPE_FRAME_ACKNOWLEDGE = 0x001E
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240549.aspx
*/
type MajorType uint16
const /*MajorType*/ (
OSMAJORTYPE_UNSPECIFIED MajorType = 0x0000
OSMAJORTYPE_WINDOWS = 0x0001
OSMAJORTYPE_OS2 = 0x0002
OSMAJORTYPE_MACINTOSH = 0x0003
OSMAJORTYPE_UNIX = 0x0004
OSMAJORTYPE_IOS = 0x0005
OSMAJORTYPE_OSX = 0x0006
OSMAJORTYPE_ANDROID = 0x0007
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240549.aspx
*/
type MinorType uint16
const /*MinorType*/ (
OSMINORTYPE_UNSPECIFIED MinorType = 0x0000
OSMINORTYPE_WINDOWS_31X = 0x0001
OSMINORTYPE_WINDOWS_95 = 0x0002
OSMINORTYPE_WINDOWS_NT = 0x0003
OSMINORTYPE_OS2_V21 = 0x0004
OSMINORTYPE_POWER_PC = 0x0005
OSMINORTYPE_MACINTOSH = 0x0006
OSMINORTYPE_NATIVE_XSERVER = 0x0007
OSMINORTYPE_PSEUDO_XSERVER = 0x0008
OSMINORTYPE_WINDOWS_RT = 0x0009
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240549.aspx
*/
type GeneralExtraFlag uint16
const /*GeneralExtraFlag*/ (
FASTPATH_OUTPUT_SUPPORTED GeneralExtraFlag = 0x0001
NO_BITMAP_COMPRESSION_HDR = 0x0400
LONG_CREDENTIALS_SUPPORTED = 0x0004
AUTORECONNECT_SUPPORTED = 0x0008
ENC_SALTED_CHECKSUM = 0x0010
)

type Boolean uint8
const /*Boolean*/ (
FALSE Boolean = 0x00
TRUE = 0x01
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240556.aspx
*/
type OrderFlag uint16
const /*OrderFlag*/ (
NEGOTIATEORDERSUPPORT OrderFlag = 0x0002
ZEROBOUNDSDELTASSUPPORT = 0x0008
COLORINDEXSUPPORT = 0x0020
SOLIDPATTERNBRUSHONLY = 0x0040
ORDERFLAGS_EXTRA_FLAGS = 0x0080
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240556.aspx
*/
type Order uint8
const /*Order*/ (
TS_NEG_DSTBLT_INDEX = 0x00
TS_NEG_PATBLT_INDEX = 0x01
TS_NEG_SCRBLT_INDEX = 0x02
TS_NEG_MEMBLT_INDEX = 0x03
TS_NEG_MEM3BLT_INDEX = 0x04
TS_NEG_DRAWNINEGRID_INDEX = 0x07
TS_NEG_LINETO_INDEX = 0x08
TS_NEG_MULTI_DRAWNINEGRID_INDEX = 0x09
TS_NEG_SAVEBITMAP_INDEX = 0x0B
TS_NEG_MULTIDSTBLT_INDEX = 0x0F
TS_NEG_MULTIPATBLT_INDEX = 0x10
TS_NEG_MULTISCRBLT_INDEX = 0x11
TS_NEG_MULTIOPAQUERECT_INDEX = 0x12
TS_NEG_FAST_INDEX_INDEX = 0x13
TS_NEG_POLYGON_SC_INDEX = 0x14
TS_NEG_POLYGON_CB_INDEX = 0x15
TS_NEG_POLYLINE_INDEX = 0x16
TS_NEG_FAST_GLYPH_INDEX = 0x18
TS_NEG_ELLIPSE_SC_INDEX = 0x19
TS_NEG_ELLIPSE_CB_INDEX = 0x1A
TS_NEG_INDEX_INDEX = 0x1B
)

type OrderEx uint16
const /*OrderEx*/ (
ORDERFLAGS_EX_CACHE_BITMAP_REV3_SUPPORT = 0x0002
ORDERFLAGS_EX_ALTSEC_FRAME_MARKER_SUPPORT = 0x0004
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240563.aspx
*/
type InputFlags uint16
const /*InputFlags*/ (
INPUT_FLAG_SCANCODES InputFlags = 0x0001
INPUT_FLAG_MOUSEX = 0x0004
INPUT_FLAG_FASTPATH_INPUT = 0x0008
INPUT_FLAG_UNICODE = 0x0010
INPUT_FLAG_FASTPATH_INPUT2 = 0x0020
INPUT_FLAG_UNUSED1 = 0x0040
INPUT_FLAG_UNUSED2 = 0x0080
TS_INPUT_FLAG_MOUSE_HWHEEL = 0x0100
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240564.aspx
*/
type BrushSupport uint32
const /*BrushSupport*/ (
BRUSH_DEFAULT BrushSupport = 0x00000000
BRUSH_COLOR_8x8 = 0x00000001
BRUSH_COLOR_FULL = 0x00000002
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240565.aspx
*/
type GlyphSupport uint16
const /*GlyphSupport*/ (
GLYPH_SUPPORT_NONE GlyphSupport = 0x0000
GLYPH_SUPPORT_PARTIAL = 0x0001
GLYPH_SUPPORT_FULL = 0x0002
GLYPH_SUPPORT_ENCODE = 0x0003
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240550.aspx
*/
type OffscreenSupportLevel uint32

const /*OffscreenSupportLevel*/ (
OSL_FALSE = 0x00000000
OSL_TRUE = 0x00000001
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240551.aspx
*/
type VirtualChannelCompressionFlag uint32
const /*VirtualChannelCompressionFlag*/ (
VCCAPS_NO_COMPR VirtualChannelCompressionFlag = 0x00000000
VCCAPS_COMPR_SC = 0x00000001
VCCAPS_COMPR_CS_8K = 0x00000002
)

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240552.aspx
*/
type SoundFlag uint16
const /*SoundFlag*/ (
SOUND_NONE = 0x0000
SOUND_BEEPS_FLAG = 0x0001
)

type Capability interface {
	GetType() CapsType
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240549.aspx
 * @param opt {object} type options
 * @returns {type.Component}
 */

type GeneralCapability struct {
	OsMajorType uint16
	OsMinorType uint16
	ProtocolVersion uint16 "0x0200, {constant : true})'"
	//pad uint16
	GeneralCompressionTypes uint16 "(0, {constant : true})"
	ExtraFlags uint16
	UpdateCapabilityFlag uint16 "(0, {constant : true})"
	RemoteUnshareFlag uint16 "(0, {constant : true})"
	GeneralCompressionLevel uint16 "(0, {constant : true})"
	RefreshRectSupport uint8
	SuppressOutputSupport uint8
}

func NewGeneralCapability() *GeneralCapability {
	return &GeneralCapability{0,0,
	0x0200, 0, 0, 0, 0,
	0, 0, 0}
}

func (caps *GeneralCapability) GetType() CapsType {
	 return CAPSTYPE_GENERAL
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240554.aspx
 * @param opt {object} type options
 * @returns {type.Component}
 */
type  BitmapCapability struct {
	preferredBitsPerPixel uint16
	receive1BitPerPixel uint16
	receive4BitsPerPixel uint16
	receive8BitsPerPixel uint16
	desktopWidth uint16
	desktopHeight uint16
	// pad2octets uint16
	desktopResizeFlag uint16
	bitmapCompressionFlag uint16
	highColorFlags uint8
	drawingFlags uint8
	multipleRectangleSupport uint16
	// pad2octetsB uint16
}

func NewBitmapCapability() *BitmapCapability {
	return &BitmapCapability{ 0,
	0x0001, 0x0001, 0x0001, 0, 0,
	0, 0x0001, 0, 0, 0x0001}
}

func (caps *GeneralCapability) GetType() CapsType {
	 return CAPSTYPE_BITMAP
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240556.aspx
 * @param orders {type.BinaryString|null} list of available orders
 * @param opt {object} type options
 * @returns {type.Component}
*/
type OrderCapability struct {
	terminalDescriptor [16]byte
	// pad4octetsA uint32
	desktopSaveXGranularity uint16
	desktopSaveYGranularity uint16
	// pad2octetsA uint16
	maximumOrderLevel uint16
	numberFonts uint16
	orderFlags OrderFlag
	orderSupport [32]byte
	textFlags uint16
	orderSupportExFlags uint16
	pad4octetsB uint32
	desktopSaveSize uint32
	// pad2octetsC uint16
	// pad2octetsD uint16
	textANSICodePage uint16
	// pad2octetsE uint16
}

func NewOrderCapability(orders *[32]byte) *OrderCapability {
	if orders != nil && len(orders) != 32 {
		panic ("NODE_RDP_PROTOCOL_PDU_CAPS_BAD_ORDERS_SIZE")
	}
	r := &OrderCapability{desktopSaveXGranularity: 1, desktopSaveYGranularity: 20, maximumOrderLevel:1,
		orderFlags: NEGOTIATEORDERSUPPORT, desktopSaveSize: 480 * 480}
	r.orderSupport = *orders
	return r
}

func (caps *OrderCapability) GetType() CapsType {
	return CAPSTYPE_ORDER
}

/**

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240559.aspx
 * @param opt type options
 * @returns {type.Component}
*/
type BitmapCacheCapability struct {
	pad1 uint32
	pad2 uint32
	pad3 uint32
	pad4 uint32
	pad5 uint32
	pad6 uint32
	cache0Entries uint16
	cache0MaximumCellSize uint16
	cache1Entries uint16
	cache1MaximumCellSize uint16
	cache2Entries uint16
	cache2MaximumCellSize uint16
}



func NewBitmapCacheCapability() *BitmapCacheCapability {
	return &BitmapCacheCapability{}
}

func (caps *OrderCapability) GetType() CapsType {
	return CAPSTYPE_BITMAPCACHE
}

/**
 *
 * @param isServer {boolean} true if in server mode
 * @param opt {object} type options
 * @returns {type.Component}
 */
type  PointerCapability struct {
	colorPointerFlag uint16
	colorPointerCacheSize uint16
	//old version of rdp doesn't support ...
	pointerCacheSize uint16
};

func NewPointerCapability(isServer bool) *PointerCapability {

	return &PointerCapability{ 0, 20,
//old version of rdp doesn't support ...
0 /*{conditional : function() { return isServer || false; }}*/}
}

func (caps *OrderCapability) GetType() CapsType {
	return CAPSTYPE_POINTER
}


/**
 * @see http://msdn.microsoft.com/en-us/library/cc240563.aspx
 * @param opt {object} type options
 * @returns {type.Component}
*/

type InputCapability struct {
	inputFlags uint16
	//pad2octetsA uint16
	// same value as gcc.ClientCoreSettings.kbdLayout
	keyboardLayout uint32
	// same value as gcc.ClientCoreSettings.keyboardType
	keyboardType uint32
	// same value as gcc.ClientCoreSettings.keyboardSubType
	keyboardSubType uint32
	// same value as gcc.ClientCoreSettings.keyboardFnKeys
	keyboardFunctionKey uint32
	// same value as gcc.ClientCoreSettingrrs.imeFileName
	imeFileName [64]byte
};


func NewInputCapability() *InputCapability {
	return &PointerCapability{}
}

func (caps *InputCapability) GetType() CapsType {
	return CAPSTYPE_INPUT
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240564.aspx
 * @param opt {object} type options
 * @returns {type.Component}
*/
type BrushCapability struct {
	brushSupportLevel BrushSupport
}
func NewBrushCapability() *BrushCapability {
	return &BrushCapability{BRUSH_DEFAULT}
}


func (caps *BrushCapability) GetType() CapsType {
	return CAPSTYPE_BRUSH
}


/**
 * @see http://msdn.microsoft.com/en-us/library/cc240566.aspx
 * @param opt {object} type options
 * @returns {type.Component}
*/
type cacheEntry struct {
	cacheEntries uint16
	cacheMaximumCellSize uint16
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240565.aspx
 * @param entries {type.Component} cache entries
 * @param opt {object} type options
 * @returns {type.Component}
 */
/*function glyphCapability(entries, opt) {
var self = {
__TYPE__ : CAPSTYPE_GLYPHCACHE,
glyphCache : entries || new type.Factory(function(s) {
c.glyphCache = new type.Component([]);
for(var i = 0; i < 10; i++) {
c.glyphCache.obj.push(cacheEntry().read(s));
}
}),
fragCache uint32
// all fonts are sent with bitmap format (very expensive)
glyphSupportLevel uint16 "(GlyphSupport.GLYPH_SUPPORT_NONE),
pad2octets uint16 "()
};

return new type.Component(self, opt);
}
*/

type GlyphCapability struct {
	GlyphCache [10] cacheEntry
	FragCache uint32
	GlyphSupportLevel GlyphSupport
	// pad2octets uint16
}

func NewGlyphCapability(entries *[10] cacheEntry) *GlyphCapability{
	c := &GlyphCapability{ GlyphSupportLevel: GLYPH_SUPPORT_NONE}
	if entries != nil {
		c.GlyphCache = *entries
	}
	return c;
}

func (caps *GlyphCapability) GetType() CapsType {
	return CAPSTYPE_GLYPHCACHE
}
/**
 * @see http://msdn.microsoft.com/en-us/library/cc240550.aspx
 * @param opt {object} type options
 * @returns {type.Component}
*/
type OffscreenBitmapCacheCapability struct {
	OffscreenSupportLevel OffscreenSupportLevel
	OffscreenCacheSize    uint16
	OffscreenCacheEntries uint16
}

func NewOffscreenBitmapCacheCapability() *OffscreenBitmapCacheCapability {
	return &OffscreenBitmapCacheCapability{ OffscreenSupportLevel: OSL_FALSE}
}

func (caps *OffscreenBitmapCacheCapability) GetType() CapsType {
	return CAPSTYPE_OFFSCREENCACHE
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240551.aspx
 * @param opt {object} type options
 * @returns {type.Component}
*/
type VirtualChannelCapability struct {
	Flags VirtualChannelCompressionFlag
	VCChunkSize uint32 "optional"
}

func NewVirtualChannelCapability() *VirtualChannelCapability {
	return &VirtualChannelCapability{ VCCAPS_NO_COMPR, 0}
}
func (caps *VirtualChannelCapability) GetType() CapsType {
	return CAPSTYPE_VIRTUALCHANNEL
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240552.aspx
 * @param opt {object} type options
 * @returns {type.Component}
 */
type SoundCapability struct {
	soundFlags SoundFlag
	// pad2octetsA uint16
}

func NewSoundCapability() *SoundCapability {
	return &SoundCapability{SOUND_NONE}
}

func (caps *SoundCapability) GetType() CapsType {
	return CAPSTYPE_SOUND
}
/**
 * @see http://msdn.microsoft.com/en-us/library/cc240568.aspx
 * @param opt {object} type options
 * @returns {type.Component}
*/
type ControlCapability struct {
	ControlFlags uint16
	RemoteDetachFlag uint16
	ControlInterest uint16
	DetachInterest uint16
}


func NewControlCapability() *ControlCapability {
	return &ControlCapability{0,
	0,0x0002, 0x0002}
}

func (caps *ControlCapability) GetType() CapsType {
	return CAPSTYPE_CONTROL
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240569.aspx
 * @param opt {object} type options
 * @returns {type.Component}
 */
type WindowActivationCapability struct {
	helpKeyFlag uint16
	helpKeyIndexFlag uint16
	helpExtendedKeyFlag uint16
	windowManagerKeyFlag uint16
}
func NewWindowActivationCapability() *WindowActivationCapability {
	return &WindowActivationCapability{}
}

func (caps *WindowActivationCapability) GetType() CapsType {
	return CAPSTYPE_ACTIVATION
}
/**
 * @see http://msdn.microsoft.com/en-us/library/cc240571.aspx
 * @param opt {object} type options
 * @returns {type.Component}
 */
type FontCapability struct {
	FontSupportFlags uint16
	// pad2octets uint16
}

func NewFontCapability() *FontCapability {
	return &FontCapability{ 0x0001}
}

func (caps *WindowActivationCapability) GetType() CapsType {
	return CAPSTYPE_FONT
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc241564.aspx
 * @param opt {object} type options
 * @returns {type.Component}
 */
type ColorCacheCapability struct {
	colorTableCacheSize uint16
	// pad2octets uint16
}

func NewColorCacheCapability() *ColorCacheCapability {
	return &ColorCacheCapability{0x0006}
}

func (caps *ColorCacheCapability) GetType() CapsType {
	return CAPSTYPE_COLORCACHE
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240570.aspx
 * @param opt {object} type options
 * @returns {type.Component}
*/
type ShareCapability struct {
	NodeId uint16
	// pad2octets uint16
}

func NewShareCapability() *ShareCapability {
	return &ShareCapability{0}
}


func (caps *ColorCacheCapability) GetType() CapsType {
	return CAPSTYPE_SHARE
}

/**
 * @see http://msdn.microsoft.com/en-us/library/cc240649.aspx
 * @param opt {object} type options
 * @returns {type.Component}
 */
type MultiFragmentUpdate struct {
	MaxRequestSize uint32
}

func NewMultiFragmentUpdate() *MultiFragmentUpdate {
	return &MultiFragmentUpdate{0}
}

func (caps *ColorCacheCapability) GetType() CapsType {
	return CAPSETTYPE_MULTIFRAGMENTUPDATE
}

/**
 * Capability wrapper packet
 * @see http://msdn.microsoft.com/en-us/library/cc240486.aspx
 * @param cap {type.Component}
 * @param opt {object} type options
 * @returns {type.Component}
 */
type CapabilityWrapper struct {
	CapabilitySetType uint16
	LengthCapability uint16
	Capability interface{}
}

func NewCapabiltyWrapper() *CapabilityWrapper {
	return &CapabilityWrapper{}
}

func (c *CapabilityWrapper) Write(writer *core.Writer) error{
	return errors.New("Нереализовано")
}

func (c *CapabilityWrapper) Read(reader *core.Reader) error {
	var err error
	c.CapabilitySetType, err = core.ReadUInt16LE(reader)
	if err != nil {
		return err
	}
	c.LengthCapability, err = core.ReadUInt16LE(reader)
	limitedReader := core.NewLimitedReader(reader, int(c.LengthCapability - 4))
	switch CapsType(c.CapabilitySetType) {
	case CAPSTYPE_GENERAL:
		c.Capability = NewGeneralCapability()
	case CAPSTYPE_BITMAP:
		c.Capability = NewBitmapCapability()
	case CAPSTYPE_ORDER:
		c.Capability = NewOrderCapability(nil)
	case CAPSTYPE_BITMAPCACHE:
		c.Capability = NewBitmapCacheCapability()
	case CAPSTYPE_POINTER:
		c.Capability = NewPointerCapability(false)
	case CAPSTYPE_INPUT:
		c.Capability = NewInputCapability()
	case CAPSTYPE_BRUSH:
		c.Capability = NewBrushCapability()
	case CAPSTYPE_GLYPHCACHE:
		c.Capability = NewGlyphCapability(nil)
	case CAPSTYPE_OFFSCREENCACHE:
		c.Capability = NewOffscreenBitmapCacheCapability()
	case CAPSTYPE_VIRTUALCHANNEL:
		c.Capability = NewVirtualChannelCapability()
	case CAPSTYPE_SOUND:
		c.Capability = NewSoundCapability()
	case CAPSTYPE_CONTROL:
		c.Capability = NewControlCapability()
	case CAPSTYPE_ACTIVATION:
		c.Capability = NewWindowActivationCapability()
	case CAPSTYPE_FONT:
		c.Capability = NewFontCapability()
	case CAPSTYPE_COLORCACHE:
		c.Capability = NewColorCacheCapability()
	case CAPSTYPE_SHARE:
		c.Capability = NewShareCapability()
	case CAPSETTYPE_MULTIFRAGMENTUPDATE:
		c.Capability = NewMultiFragmentUpdate()

	default:
		log.Printf("unknown capability %V", c.CapabilitySetType)
		c.Capability, _ = core.AllocAndReadBytes(int(c.LengthCapability - 4), reader)
		return nil
	}
	return c.Capability.(core.Readable).Read(limitedReader)
}



