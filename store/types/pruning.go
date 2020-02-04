/*
BugFix: extend with helpers related CommitInfo flush bug
Copyright (C) 2020 Ethan Frey

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package types

// PruningStrategy specifies how old states will be deleted over time where
// keepRecent can be used with keepEvery to create a pruning "strategy".
type PruningOptions struct {
	keepRecent int64
	keepEvery  int64
}

func NewPruningOptions(keepRecent, keepEvery int64) PruningOptions {
	return PruningOptions{
		keepRecent: keepRecent,
		keepEvery:  keepEvery,
	}
}

// How much recent state will be kept. Older state will be deleted.
func (po PruningOptions) KeepRecent() int64 {
	return po.keepRecent
}

// Keeps every N stated, deleting others.
func (po PruningOptions) KeepEvery() int64 {
	return po.keepEvery
}

// WillFlush returns true if this height is writen to disk
func (po PruningOptions) WillFlush(height int64) bool {
	return height%po.keepEvery == 0
}

// default pruning strategies
var (
	// PruneEverything means all saved states will be deleted, storing only the current state
	PruneEverything = NewPruningOptions(1, 1)
	// PruneNothing means all historic states will be saved, nothing will be deleted
	PruneNothing = NewPruningOptions(0, 1)
	// PruneSyncable means only those states not needed for state syncing will be deleted (keeps last 100 + every 10000th)
	// TODO: update, only write every 10 for now
	PruneSyncable = NewPruningOptions(10, 10)
)
