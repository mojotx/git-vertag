package internal

type Manager struct {
	Tagger Tagger
}

func (m *Manager) DeleteVer(v Semver) error {
	if err := m.Tagger.DeleteTag(v.String()); err != nil {
		return err
	}
	return nil
}

func (m *Manager) CreateVer(v Semver, message []string, file string) error {
	if err := m.Tagger.CreateTag(v.String(), message, file); err != nil {
		return err
	}
	return nil
}

func (m *Manager) ReplaceVer(v Semver, message []string, file string) error {
	if err := m.DeleteVer(v); err != nil {
		return err
	}
	if err := m.CreateVer(v, message, file); err != nil {
		return err
	}
	return nil
}

func (m *Manager) GetVer(fetch bool) (Semver, error) {
	var latest Semver
	tags, err := m.Tagger.GetTags(fetch)
	if err != nil {
		return latest, err
	}
	for _, tag := range tags {
		ver, err := Parse(tag)
		if err != nil {
			continue
		}
		latest = GreaterSemver(latest, ver)
	}

	return latest, nil
}