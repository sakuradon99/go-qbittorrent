package model

type GetTorrentListOptions struct {
	// Filter torrent list by state
	Filter TorrentListFilter `url:"filter,omitempty"`
	// Category Get torrents with the given category (empty string means "without category"; no "category" parameter
	//means "any category" broken until #11748 is resolved). Remember to URL-encode the category name. For example,
	//My category becomes My%20category
	Category *string `url:"category,omitempty"`
	// Tag Get torrents with the given tag (empty string means "without tag"; no "tag" parameter means "any tag".)
	// Remember to URL-encode the category name. For example, My tag becomes My%20tag
	Tag string `url:"tag,omitempty"`
	// Sort torrents by given key. They can be sorted using any field of the response's JSON array (which are
	//documented below) as the sort key.
	Sort string `url:"sort,omitempty"`
	// Reverse Enable reverse sorting. Defaults to false
	Reverse bool `url:"reverse,omitempty"`
	// Limit the number of torrents returned
	Limit int `url:"limit,omitempty"`
	// Offset Set offset (if less than 0, offset from end)
	Offset int `url:"offset,omitempty"`
	// Hashes Filter by hashes. Can contain multiple hashes separated by |
	Hashes string `url:"hashes,omitempty"`
}

type TorrentListFilter string

const (
	FilterAll         TorrentListFilter = "all"
	FilterDownloading TorrentListFilter = "downloading"
	FilterSeeding     TorrentListFilter = "seeding"
	FilterCompleted   TorrentListFilter = "completed"
	FilterPaused      TorrentListFilter = "paused"
	FilterActive      TorrentListFilter = "active"
	FilterInactive    TorrentListFilter = "inactive"
	FilterResumed     TorrentListFilter = "resumed"
	FilterStalled     TorrentListFilter = "stalled"
	FilterStalledUp   TorrentListFilter = "stalled_uploading"
	FilerStalledDl    TorrentListFilter = "stalled_downloading"
	FilterErrored     TorrentListFilter = "errored"
)
