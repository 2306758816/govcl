
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
    "time"
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TMonthCalendar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewMonthCalendar
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMonthCalendar(owner IComponent) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = MonthCalendar_Create(CheckPtr(owner))
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MonthCalendarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func MonthCalendarFromInst(inst uintptr) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = inst
    m.ptr = unsafe.Pointer(inst)
    return m
}

// MonthCalendarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func MonthCalendarFromObj(obj IObject) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = CheckPtr(obj)
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MonthCalendarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func MonthCalendarFromUnsafePointer(ptr unsafe.Pointer) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = uintptr(ptr)
    m.ptr = ptr
    return m
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (m *TMonthCalendar) Free() {
    if m.instance != 0 {
        MonthCalendar_Free(m.instance)
        m.instance = 0
        m.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMonthCalendar) Instance() uintptr {
    return m.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMonthCalendar) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMonthCalendar) IsValid() bool {
    return m.instance != 0
}

// TMonthCalendarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMonthCalendarClass() TClass {
    return MonthCalendar_StaticClassType()
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (m *TMonthCalendar) CanFocus() bool {
    return MonthCalendar_CanFocus(m.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (m *TMonthCalendar) ContainsControl(Control IControl) bool {
    return MonthCalendar_ContainsControl(m.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (m *TMonthCalendar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(MonthCalendar_ControlAtPos(m.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (m *TMonthCalendar) DisableAlign() {
    MonthCalendar_DisableAlign(m.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (m *TMonthCalendar) EnableAlign() {
    MonthCalendar_EnableAlign(m.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (m *TMonthCalendar) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(MonthCalendar_FindChildControl(m.instance, ControlName))
}

// FlipChildren
func (m *TMonthCalendar) FlipChildren(AllLevels bool) {
    MonthCalendar_FlipChildren(m.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (m *TMonthCalendar) Focused() bool {
    return MonthCalendar_Focused(m.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (m *TMonthCalendar) HandleAllocated() bool {
    return MonthCalendar_HandleAllocated(m.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (m *TMonthCalendar) InsertControl(AControl IControl) {
    MonthCalendar_InsertControl(m.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (m *TMonthCalendar) Invalidate() {
    MonthCalendar_Invalidate(m.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (m *TMonthCalendar) PaintTo(DC HDC, X int32, Y int32) {
    MonthCalendar_PaintTo(m.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (m *TMonthCalendar) RemoveControl(AControl IControl) {
    MonthCalendar_RemoveControl(m.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (m *TMonthCalendar) Realign() {
    MonthCalendar_Realign(m.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (m *TMonthCalendar) Repaint() {
    MonthCalendar_Repaint(m.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (m *TMonthCalendar) ScaleBy(M int32, D int32) {
    MonthCalendar_ScaleBy(m.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (m *TMonthCalendar) ScrollBy(DeltaX int32, DeltaY int32) {
    MonthCalendar_ScrollBy(m.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (m *TMonthCalendar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    MonthCalendar_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (m *TMonthCalendar) SetFocus() {
    MonthCalendar_SetFocus(m.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (m *TMonthCalendar) Update() {
    MonthCalendar_Update(m.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (m *TMonthCalendar) UpdateControlState() {
    MonthCalendar_UpdateControlState(m.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (m *TMonthCalendar) BringToFront() {
    MonthCalendar_BringToFront(m.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (m *TMonthCalendar) ClientToScreen(Point TPoint) TPoint {
    return MonthCalendar_ClientToScreen(m.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (m *TMonthCalendar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return MonthCalendar_ClientToParent(m.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (m *TMonthCalendar) Dragging() bool {
    return MonthCalendar_Dragging(m.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (m *TMonthCalendar) HasParent() bool {
    return MonthCalendar_HasParent(m.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (m *TMonthCalendar) Hide() {
    MonthCalendar_Hide(m.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (m *TMonthCalendar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return MonthCalendar_Perform(m.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (m *TMonthCalendar) Refresh() {
    MonthCalendar_Refresh(m.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (m *TMonthCalendar) ScreenToClient(Point TPoint) TPoint {
    return MonthCalendar_ScreenToClient(m.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (m *TMonthCalendar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return MonthCalendar_ParentToClient(m.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (m *TMonthCalendar) SendToBack() {
    MonthCalendar_SendToBack(m.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (m *TMonthCalendar) Show() {
    MonthCalendar_Show(m.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (m *TMonthCalendar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return MonthCalendar_GetTextBuf(m.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (m *TMonthCalendar) GetTextLen() int32 {
    return MonthCalendar_GetTextLen(m.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (m *TMonthCalendar) SetTextBuf(Buffer string) {
    MonthCalendar_SetTextBuf(m.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (m *TMonthCalendar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(MonthCalendar_FindComponent(m.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (m *TMonthCalendar) GetNamePath() string {
    return MonthCalendar_GetNamePath(m.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (m *TMonthCalendar) Assign(Source IObject) {
    MonthCalendar_Assign(m.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (m *TMonthCalendar) DisposeOf() {
    MonthCalendar_DisposeOf(m.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMonthCalendar) ClassType() TClass {
    return MonthCalendar_ClassType(m.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMonthCalendar) ClassName() string {
    return MonthCalendar_ClassName(m.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMonthCalendar) InstanceSize() int32 {
    return MonthCalendar_InstanceSize(m.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMonthCalendar) InheritsFrom(AClass TClass) bool {
    return MonthCalendar_InheritsFrom(m.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMonthCalendar) Equals(Obj IObject) bool {
    return MonthCalendar_Equals(m.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMonthCalendar) GetHashCode() int32 {
    return MonthCalendar_GetHashCode(m.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (m *TMonthCalendar) ToString() string {
    return MonthCalendar_ToString(m.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (m *TMonthCalendar) Align() TAlign {
    return MonthCalendar_GetAlign(m.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (m *TMonthCalendar) SetAlign(value TAlign) {
    MonthCalendar_SetAlign(m.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (m *TMonthCalendar) Anchors() TAnchors {
    return MonthCalendar_GetAnchors(m.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (m *TMonthCalendar) SetAnchors(value TAnchors) {
    MonthCalendar_SetAnchors(m.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (m *TMonthCalendar) AutoSize() bool {
    return MonthCalendar_GetAutoSize(m.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
func (m *TMonthCalendar) SetAutoSize(value bool) {
    MonthCalendar_SetAutoSize(m.instance, value)
}

// BorderWidth
// CN: 获取边框的宽度。
// EN: .
func (m *TMonthCalendar) BorderWidth() int32 {
    return MonthCalendar_GetBorderWidth(m.instance)
}

// SetBorderWidth
// CN: 设置边框的宽度。
// EN: .
func (m *TMonthCalendar) SetBorderWidth(value int32) {
    MonthCalendar_SetBorderWidth(m.instance, value)
}

// BiDiMode
func (m *TMonthCalendar) BiDiMode() TBiDiMode {
    return MonthCalendar_GetBiDiMode(m.instance)
}

// SetBiDiMode
func (m *TMonthCalendar) SetBiDiMode(value TBiDiMode) {
    MonthCalendar_SetBiDiMode(m.instance, value)
}

// CalColors
func (m *TMonthCalendar) CalColors() *TMonthCalColors {
    return MonthCalColorsFromInst(MonthCalendar_GetCalColors(m.instance))
}

// SetCalColors
func (m *TMonthCalendar) SetCalColors(value *TMonthCalColors) {
    MonthCalendar_SetCalColors(m.instance, CheckPtr(value))
}

// MultiSelect
func (m *TMonthCalendar) MultiSelect() bool {
    return MonthCalendar_GetMultiSelect(m.instance)
}

// SetMultiSelect
func (m *TMonthCalendar) SetMultiSelect(value bool) {
    MonthCalendar_SetMultiSelect(m.instance, value)
}

// Date
func (m *TMonthCalendar) Date() time.Time {
    return MonthCalendar_GetDate(m.instance)
}

// SetDate
func (m *TMonthCalendar) SetDate(value time.Time) {
    MonthCalendar_SetDate(m.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (m *TMonthCalendar) DoubleBuffered() bool {
    return MonthCalendar_GetDoubleBuffered(m.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (m *TMonthCalendar) SetDoubleBuffered(value bool) {
    MonthCalendar_SetDoubleBuffered(m.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (m *TMonthCalendar) DragCursor() TCursor {
    return MonthCalendar_GetDragCursor(m.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (m *TMonthCalendar) SetDragCursor(value TCursor) {
    MonthCalendar_SetDragCursor(m.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (m *TMonthCalendar) DragKind() TDragKind {
    return MonthCalendar_GetDragKind(m.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (m *TMonthCalendar) SetDragKind(value TDragKind) {
    MonthCalendar_SetDragKind(m.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (m *TMonthCalendar) DragMode() TDragMode {
    return MonthCalendar_GetDragMode(m.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (m *TMonthCalendar) SetDragMode(value TDragMode) {
    MonthCalendar_SetDragMode(m.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (m *TMonthCalendar) Enabled() bool {
    return MonthCalendar_GetEnabled(m.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (m *TMonthCalendar) SetEnabled(value bool) {
    MonthCalendar_SetEnabled(m.instance, value)
}

// FirstDayOfWeek
func (m *TMonthCalendar) FirstDayOfWeek() TCalDayOfWeek {
    return MonthCalendar_GetFirstDayOfWeek(m.instance)
}

// SetFirstDayOfWeek
func (m *TMonthCalendar) SetFirstDayOfWeek(value TCalDayOfWeek) {
    MonthCalendar_SetFirstDayOfWeek(m.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (m *TMonthCalendar) Font() *TFont {
    return FontFromInst(MonthCalendar_GetFont(m.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (m *TMonthCalendar) SetFont(value *TFont) {
    MonthCalendar_SetFont(m.instance, CheckPtr(value))
}

// MaxDate
func (m *TMonthCalendar) MaxDate() time.Time {
    return MonthCalendar_GetMaxDate(m.instance)
}

// SetMaxDate
func (m *TMonthCalendar) SetMaxDate(value time.Time) {
    MonthCalendar_SetMaxDate(m.instance, value)
}

// MaxSelectRange
func (m *TMonthCalendar) MaxSelectRange() int32 {
    return MonthCalendar_GetMaxSelectRange(m.instance)
}

// SetMaxSelectRange
func (m *TMonthCalendar) SetMaxSelectRange(value int32) {
    MonthCalendar_SetMaxSelectRange(m.instance, value)
}

// MinDate
func (m *TMonthCalendar) MinDate() time.Time {
    return MonthCalendar_GetMinDate(m.instance)
}

// SetMinDate
func (m *TMonthCalendar) SetMinDate(value time.Time) {
    MonthCalendar_SetMinDate(m.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (m *TMonthCalendar) ParentDoubleBuffered() bool {
    return MonthCalendar_GetParentDoubleBuffered(m.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (m *TMonthCalendar) SetParentDoubleBuffered(value bool) {
    MonthCalendar_SetParentDoubleBuffered(m.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (m *TMonthCalendar) ParentFont() bool {
    return MonthCalendar_GetParentFont(m.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (m *TMonthCalendar) SetParentFont(value bool) {
    MonthCalendar_SetParentFont(m.instance, value)
}

// ParentShowHint
func (m *TMonthCalendar) ParentShowHint() bool {
    return MonthCalendar_GetParentShowHint(m.instance)
}

// SetParentShowHint
func (m *TMonthCalendar) SetParentShowHint(value bool) {
    MonthCalendar_SetParentShowHint(m.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (m *TMonthCalendar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(MonthCalendar_GetPopupMenu(m.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (m *TMonthCalendar) SetPopupMenu(value IComponent) {
    MonthCalendar_SetPopupMenu(m.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (m *TMonthCalendar) ShowHint() bool {
    return MonthCalendar_GetShowHint(m.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (m *TMonthCalendar) SetShowHint(value bool) {
    MonthCalendar_SetShowHint(m.instance, value)
}

// ShowToday
func (m *TMonthCalendar) ShowToday() bool {
    return MonthCalendar_GetShowToday(m.instance)
}

// SetShowToday
func (m *TMonthCalendar) SetShowToday(value bool) {
    MonthCalendar_SetShowToday(m.instance, value)
}

// ShowTodayCircle
func (m *TMonthCalendar) ShowTodayCircle() bool {
    return MonthCalendar_GetShowTodayCircle(m.instance)
}

// SetShowTodayCircle
func (m *TMonthCalendar) SetShowTodayCircle(value bool) {
    MonthCalendar_SetShowTodayCircle(m.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (m *TMonthCalendar) TabOrder() TTabOrder {
    return MonthCalendar_GetTabOrder(m.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (m *TMonthCalendar) SetTabOrder(value TTabOrder) {
    MonthCalendar_SetTabOrder(m.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (m *TMonthCalendar) TabStop() bool {
    return MonthCalendar_GetTabStop(m.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (m *TMonthCalendar) SetTabStop(value bool) {
    MonthCalendar_SetTabStop(m.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (m *TMonthCalendar) Visible() bool {
    return MonthCalendar_GetVisible(m.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (m *TMonthCalendar) SetVisible(value bool) {
    MonthCalendar_SetVisible(m.instance, value)
}

// WeekNumbers
func (m *TMonthCalendar) WeekNumbers() bool {
    return MonthCalendar_GetWeekNumbers(m.instance)
}

// SetWeekNumbers
func (m *TMonthCalendar) SetWeekNumbers(value bool) {
    MonthCalendar_SetWeekNumbers(m.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (m *TMonthCalendar) SetOnClick(fn TNotifyEvent) {
    MonthCalendar_SetOnClick(m.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (m *TMonthCalendar) SetOnContextPopup(fn TContextPopupEvent) {
    MonthCalendar_SetOnContextPopup(m.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (m *TMonthCalendar) SetOnDblClick(fn TNotifyEvent) {
    MonthCalendar_SetOnDblClick(m.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (m *TMonthCalendar) SetOnDragDrop(fn TDragDropEvent) {
    MonthCalendar_SetOnDragDrop(m.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (m *TMonthCalendar) SetOnDragOver(fn TDragOverEvent) {
    MonthCalendar_SetOnDragOver(m.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (m *TMonthCalendar) SetOnEndDock(fn TEndDragEvent) {
    MonthCalendar_SetOnEndDock(m.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (m *TMonthCalendar) SetOnEndDrag(fn TEndDragEvent) {
    MonthCalendar_SetOnEndDrag(m.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (m *TMonthCalendar) SetOnEnter(fn TNotifyEvent) {
    MonthCalendar_SetOnEnter(m.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (m *TMonthCalendar) SetOnExit(fn TNotifyEvent) {
    MonthCalendar_SetOnExit(m.instance, fn)
}

// SetOnGesture
func (m *TMonthCalendar) SetOnGesture(fn TGestureEvent) {
    MonthCalendar_SetOnGesture(m.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (m *TMonthCalendar) SetOnKeyDown(fn TKeyEvent) {
    MonthCalendar_SetOnKeyDown(m.instance, fn)
}

// SetOnKeyPress
func (m *TMonthCalendar) SetOnKeyPress(fn TKeyPressEvent) {
    MonthCalendar_SetOnKeyPress(m.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (m *TMonthCalendar) SetOnKeyUp(fn TKeyEvent) {
    MonthCalendar_SetOnKeyUp(m.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (m *TMonthCalendar) SetOnMouseEnter(fn TNotifyEvent) {
    MonthCalendar_SetOnMouseEnter(m.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (m *TMonthCalendar) SetOnMouseLeave(fn TNotifyEvent) {
    MonthCalendar_SetOnMouseLeave(m.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (m *TMonthCalendar) SetOnStartDock(fn TStartDockEvent) {
    MonthCalendar_SetOnStartDock(m.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (m *TMonthCalendar) DockClientCount() int32 {
    return MonthCalendar_GetDockClientCount(m.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (m *TMonthCalendar) DockSite() bool {
    return MonthCalendar_GetDockSite(m.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (m *TMonthCalendar) SetDockSite(value bool) {
    MonthCalendar_SetDockSite(m.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (m *TMonthCalendar) AlignDisabled() bool {
    return MonthCalendar_GetAlignDisabled(m.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (m *TMonthCalendar) MouseInClient() bool {
    return MonthCalendar_GetMouseInClient(m.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (m *TMonthCalendar) VisibleDockClientCount() int32 {
    return MonthCalendar_GetVisibleDockClientCount(m.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (m *TMonthCalendar) Brush() *TBrush {
    return BrushFromInst(MonthCalendar_GetBrush(m.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (m *TMonthCalendar) ControlCount() int32 {
    return MonthCalendar_GetControlCount(m.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (m *TMonthCalendar) Handle() HWND {
    return MonthCalendar_GetHandle(m.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (m *TMonthCalendar) ParentWindow() HWND {
    return MonthCalendar_GetParentWindow(m.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (m *TMonthCalendar) SetParentWindow(value HWND) {
    MonthCalendar_SetParentWindow(m.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (m *TMonthCalendar) UseDockManager() bool {
    return MonthCalendar_GetUseDockManager(m.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (m *TMonthCalendar) SetUseDockManager(value bool) {
    MonthCalendar_SetUseDockManager(m.instance, value)
}

// Action
func (m *TMonthCalendar) Action() *TAction {
    return ActionFromInst(MonthCalendar_GetAction(m.instance))
}

// SetAction
func (m *TMonthCalendar) SetAction(value IComponent) {
    MonthCalendar_SetAction(m.instance, CheckPtr(value))
}

// BoundsRect
func (m *TMonthCalendar) BoundsRect() TRect {
    return MonthCalendar_GetBoundsRect(m.instance)
}

// SetBoundsRect
func (m *TMonthCalendar) SetBoundsRect(value TRect) {
    MonthCalendar_SetBoundsRect(m.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (m *TMonthCalendar) ClientHeight() int32 {
    return MonthCalendar_GetClientHeight(m.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (m *TMonthCalendar) SetClientHeight(value int32) {
    MonthCalendar_SetClientHeight(m.instance, value)
}

// ClientOrigin
func (m *TMonthCalendar) ClientOrigin() TPoint {
    return MonthCalendar_GetClientOrigin(m.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (m *TMonthCalendar) ClientRect() TRect {
    return MonthCalendar_GetClientRect(m.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (m *TMonthCalendar) ClientWidth() int32 {
    return MonthCalendar_GetClientWidth(m.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (m *TMonthCalendar) SetClientWidth(value int32) {
    MonthCalendar_SetClientWidth(m.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (m *TMonthCalendar) ControlState() TControlState {
    return MonthCalendar_GetControlState(m.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (m *TMonthCalendar) SetControlState(value TControlState) {
    MonthCalendar_SetControlState(m.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (m *TMonthCalendar) ControlStyle() TControlStyle {
    return MonthCalendar_GetControlStyle(m.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (m *TMonthCalendar) SetControlStyle(value TControlStyle) {
    MonthCalendar_SetControlStyle(m.instance, value)
}

// ExplicitLeft
func (m *TMonthCalendar) ExplicitLeft() int32 {
    return MonthCalendar_GetExplicitLeft(m.instance)
}

// ExplicitTop
func (m *TMonthCalendar) ExplicitTop() int32 {
    return MonthCalendar_GetExplicitTop(m.instance)
}

// ExplicitWidth
func (m *TMonthCalendar) ExplicitWidth() int32 {
    return MonthCalendar_GetExplicitWidth(m.instance)
}

// ExplicitHeight
func (m *TMonthCalendar) ExplicitHeight() int32 {
    return MonthCalendar_GetExplicitHeight(m.instance)
}

// Floating
func (m *TMonthCalendar) Floating() bool {
    return MonthCalendar_GetFloating(m.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (m *TMonthCalendar) Parent() *TWinControl {
    return WinControlFromInst(MonthCalendar_GetParent(m.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (m *TMonthCalendar) SetParent(value IWinControl) {
    MonthCalendar_SetParent(m.instance, CheckPtr(value))
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (m *TMonthCalendar) StyleElements() TStyleElements {
    return MonthCalendar_GetStyleElements(m.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (m *TMonthCalendar) SetStyleElements(value TStyleElements) {
    MonthCalendar_SetStyleElements(m.instance, value)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (m *TMonthCalendar) AlignWithMargins() bool {
    return MonthCalendar_GetAlignWithMargins(m.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (m *TMonthCalendar) SetAlignWithMargins(value bool) {
    MonthCalendar_SetAlignWithMargins(m.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (m *TMonthCalendar) Left() int32 {
    return MonthCalendar_GetLeft(m.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (m *TMonthCalendar) SetLeft(value int32) {
    MonthCalendar_SetLeft(m.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (m *TMonthCalendar) Top() int32 {
    return MonthCalendar_GetTop(m.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (m *TMonthCalendar) SetTop(value int32) {
    MonthCalendar_SetTop(m.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (m *TMonthCalendar) Width() int32 {
    return MonthCalendar_GetWidth(m.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (m *TMonthCalendar) SetWidth(value int32) {
    MonthCalendar_SetWidth(m.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (m *TMonthCalendar) Height() int32 {
    return MonthCalendar_GetHeight(m.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (m *TMonthCalendar) SetHeight(value int32) {
    MonthCalendar_SetHeight(m.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (m *TMonthCalendar) Cursor() TCursor {
    return MonthCalendar_GetCursor(m.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (m *TMonthCalendar) SetCursor(value TCursor) {
    MonthCalendar_SetCursor(m.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (m *TMonthCalendar) Hint() string {
    return MonthCalendar_GetHint(m.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (m *TMonthCalendar) SetHint(value string) {
    MonthCalendar_SetHint(m.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (m *TMonthCalendar) Margins() *TMargins {
    return MarginsFromInst(MonthCalendar_GetMargins(m.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (m *TMonthCalendar) SetMargins(value *TMargins) {
    MonthCalendar_SetMargins(m.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (m *TMonthCalendar) CustomHint() *TCustomHint {
    return CustomHintFromInst(MonthCalendar_GetCustomHint(m.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (m *TMonthCalendar) SetCustomHint(value IComponent) {
    MonthCalendar_SetCustomHint(m.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (m *TMonthCalendar) ComponentCount() int32 {
    return MonthCalendar_GetComponentCount(m.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (m *TMonthCalendar) ComponentIndex() int32 {
    return MonthCalendar_GetComponentIndex(m.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (m *TMonthCalendar) SetComponentIndex(value int32) {
    MonthCalendar_SetComponentIndex(m.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (m *TMonthCalendar) Owner() *TComponent {
    return ComponentFromInst(MonthCalendar_GetOwner(m.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (m *TMonthCalendar) Name() string {
    return MonthCalendar_GetName(m.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (m *TMonthCalendar) SetName(value string) {
    MonthCalendar_SetName(m.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (m *TMonthCalendar) Tag() int {
    return MonthCalendar_GetTag(m.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (m *TMonthCalendar) SetTag(value int) {
    MonthCalendar_SetTag(m.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (m *TMonthCalendar) DockClients(Index int32) *TControl {
    return ControlFromInst(MonthCalendar_GetDockClients(m.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (m *TMonthCalendar) Controls(Index int32) *TControl {
    return ControlFromInst(MonthCalendar_GetControls(m.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (m *TMonthCalendar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(MonthCalendar_GetComponents(m.instance, AIndex))
}

