// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://www.mugomes.com.br

package mgsmartflow

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

/* Meta */

type flowMeta struct {
	row  int
	size *fyne.Size
	move *fyne.Position
	gap  *fyne.Position
}

/* Layout */

type SmartFlowLayout struct {
	meta map[fyne.CanvasObject]*flowMeta

	rows [][]fyne.CanvasObject
	row  int

	gapX float32
	gapY float32
}

func NewSmartFlowLayout() *SmartFlowLayout {
	return &SmartFlowLayout{
		meta: make(map[fyne.CanvasObject]*flowMeta),
		gapX: 7,
		gapY: 7,
	}
}

/* Config */

func (l *SmartFlowLayout) SetGlobalGap(x, y float32) {
	l.gapX = x
	l.gapY = y
}

func (l *SmartFlowLayout) SetGap(obj fyne.CanvasObject, gap fyne.Position) {
	meta := l.ensureMeta(obj)
	meta.gap = &gap
}

func (l *SmartFlowLayout) SetResize(obj fyne.CanvasObject, size fyne.Size) {
	meta := l.ensureMeta(obj)
	meta.size = &size
}

func (l *SmartFlowLayout) SetMove(obj fyne.CanvasObject, pos fyne.Position) {
	meta := l.ensureMeta(obj)
	meta.move = &pos
}

/* Add */

func (l *SmartFlowLayout) AddRow(obj fyne.CanvasObject) {
	l.meta[obj] = &flowMeta{row: l.row}
	l.row++
}

func (l *SmartFlowLayout) AddColumn(objs ...fyne.CanvasObject) {
	for _, obj := range objs {
		l.meta[obj] = &flowMeta{row: l.row}
	}
	l.row++
}

/* Layout */

func (l *SmartFlowLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	l.prepareRows()

	for _, obj := range objects {
		if meta := l.meta[obj]; meta != nil {
			l.rows[meta.row] = append(l.rows[meta.row], obj)
		}
	}

	y := float32(0)

	for r := 0; r < l.row; r++ {
		row := l.rows[r]
		if len(row) == 0 {
			continue
		}

		rowHeight := float32(0)
		fixedW := float32(0)
		auto := 0

		for _, obj := range row {
			meta := l.meta[obj]
			h := obj.MinSize().Height

			if meta.size != nil {
				h = meta.size.Height
				fixedW += meta.size.Width
			} else {
				auto++
			}

			if h > rowHeight {
				rowHeight = h
			}
		}

		remain := size.Width - fixedW - l.gapX*float32(len(row)-1)
		if remain < 0 {
			remain = 0
		}

		autoW := float32(0)
		if auto > 0 {
			autoW = remain / float32(auto)
		}

		x := float32(0)

		for _, obj := range row {
			meta := l.meta[obj]

			w := autoW
			h := rowHeight

			if meta.size != nil {
				w = meta.size.Width
				h = meta.size.Height
			}

			obj.Resize(fyne.NewSize(w, h))

			if meta.move != nil {
				obj.Move(fyne.NewPos(meta.move.X, y))
			} else {
				obj.Move(fyne.NewPos(x, y))
			}

			gx := l.gapX
			if meta.gap != nil {
				gx = meta.gap.X
			}

			x += w + gx
		}

		gy := l.gapY
		for _, obj := range row {
			if meta := l.meta[obj]; meta.gap != nil && meta.gap.Y > 0 {
				gy = meta.gap.Y
				break
			}
		}

		y += rowHeight + gy
	}
}

/* MinSize */
func (l *SmartFlowLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	l.prepareRows()

	for _, obj := range objects {
		if meta := l.meta[obj]; meta != nil {
			l.rows[meta.row] = append(l.rows[meta.row], obj)
		}
	}

	totalH := float32(0)

	for r := 0; r < l.row; r++ {
		row := l.rows[r]
		if len(row) == 0 {
			continue
		}

		rowH := float32(0)

		for _, obj := range row {
			meta := l.meta[obj]
			h := obj.MinSize().Height

			if meta.size != nil {
				h = meta.size.Height
			}

			if h > rowH {
				rowH = h
			}
		}

		totalH += rowH
		if r < l.row-1 {
			totalH += l.gapY
		}
	}

	return fyne.NewSize(0, totalH)
}

/* Internals */

func (l *SmartFlowLayout) prepareRows() {
	if len(l.rows) < l.row {
		l.rows = make([][]fyne.CanvasObject, l.row)
	}
	for i := 0; i < l.row; i++ {
		l.rows[i] = l.rows[i][:0]
	}
}

func (l *SmartFlowLayout) ensureMeta(obj fyne.CanvasObject) *flowMeta {
	if l.meta[obj] == nil {
		l.meta[obj] = &flowMeta{row: l.row}
		l.row++
	}
	return l.meta[obj]
}

/* WRAPPER */

type SmartFlow struct {
	Container *fyne.Container
	Layout    *SmartFlowLayout
}

func New() *SmartFlow {
	layout := NewSmartFlowLayout()
	return &SmartFlow{
		Container: container.New(layout),
		Layout:    layout,
	}
}

/* Publico */

func (s *SmartFlow) AddRow(obj fyne.CanvasObject) {
	s.Layout.AddRow(obj)
	s.Container.Add(obj)
	s.Container.Refresh()
}

func (s *SmartFlow) AddColumn(objs ...fyne.CanvasObject) {
	s.Layout.AddColumn(objs...)
	for _, obj := range objs {
		s.Container.Add(obj)
	}
	s.Container.Refresh()
}

func (s *SmartFlow) SetResize(obj fyne.CanvasObject, size fyne.Size) {
	s.Layout.SetResize(obj, size)
	s.Container.Refresh()
}

func (s *SmartFlow) SetMove(obj fyne.CanvasObject, pos fyne.Position) {
	s.Layout.SetMove(obj, pos)
	s.Container.Refresh()
}

func (s *SmartFlow) SetGap(obj fyne.CanvasObject, gap fyne.Position) {
	s.Layout.SetGap(obj, gap)
	s.Container.Refresh()
}

func (s *SmartFlow) SetGlobalGap(x, y float32) {
	s.Layout.SetGlobalGap(x, y)
	s.Container.Refresh()
}
