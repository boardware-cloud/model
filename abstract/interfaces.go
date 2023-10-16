package abstract

type Owner interface {
	ID() uint
	Owner() Owner
	Own(asset Asset) bool
	TypeName() string
}

type Asset interface {
	ID() uint
	Owner() Owner
	SetOwner(owner Owner)
}

func Own(owner Owner, asset Owner) bool {
	if owner.TypeName() == asset.Owner().TypeName() && owner.ID() == asset.Owner().ID() {
		return true
	}
	if owner.Owner() == nil {
		return false
	}
	return Own(owner.Owner(), asset)
}
