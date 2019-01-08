package main

type Session struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type Configuration struct {
	Characters   []*Character
	Region       string
	Guild        string
	Realm        string
	GuildRanks   []int
	GuildMembers Guild
	ClientID     string
	ClientSecret string
	Session      Session
}

type Guild struct {
	LastModified      int64  `json:"lastModified"`
	Name              string `json:"name"`
	Realm             string `json:"realm"`
	Battlegroup       string `json:"battlegroup"`
	Level             int    `json:"level"`
	Side              int    `json:"side"`
	AchievementPoints int    `json:"achievementPoints"`
	Members           []struct {
		Character struct {
			Name              string `json:"name"`
			Realm             string `json:"realm"`
			Battlegroup       string `json:"battlegroup"`
			Class             int    `json:"class"`
			Race              int    `json:"race"`
			Gender            int    `json:"gender"`
			Level             int    `json:"level"`
			AchievementPoints int    `json:"achievementPoints"`
			Thumbnail         string `json:"thumbnail"`
			Spec              struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
			Guild        string `json:"guild"`
			GuildRealm   string `json:"guildRealm"`
			LastModified int    `json:"lastModified"`
		} `json:"character"`
		Rank int `json:"rank"`
	} `json:"members"`
	Emblem struct {
		Icon              int    `json:"icon"`
		IconColor         string `json:"iconColor"`
		IconColorID       int    `json:"iconColorId"`
		Border            int    `json:"border"`
		BorderColor       string `json:"borderColor"`
		BorderColorID     int    `json:"borderColorId"`
		BackgroundColor   string `json:"backgroundColor"`
		BackgroundColorID int    `json:"backgroundColorId"`
	} `json:"emblem"`
}

type Character struct {
	LastModified      int64  `json:"lastModified"`
	Name              string `json:"name"`
	Realm             string `json:"realm"`
	Battlegroup       string `json:"battlegroup"`
	Class             int    `json:"class"`
	Race              int    `json:"race"`
	Gender            int    `json:"gender"`
	Level             int    `json:"level"`
	AchievementPoints int    `json:"achievementPoints"`
	Thumbnail         string `json:"thumbnail"`
	CalcClass         string `json:"calcClass"`
	Faction           int    `json:"faction"`
	Items             struct {
		AverageItemLevel         int `json:"averageItemLevel"`
		AverageItemLevelEquipped int `json:"averageItemLevelEquipped"`
		Head                     struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem      int `json:"transmogItem"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				ItemAppearanceModID         int `json:"itemAppearanceModId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []struct {
					ID          int `json:"id"`
					Tier        int `json:"tier"`
					SpellID     int `json:"spellId"`
					BonusListID int `json:"bonusListId"`
				} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"head"`
		Neck struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"neck"`
		Shoulder struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem      int `json:"transmogItem"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				ItemAppearanceModID         int `json:"itemAppearanceModId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []struct {
					ID          int `json:"id"`
					Tier        int `json:"tier"`
					SpellID     int `json:"spellId"`
					BonusListID int `json:"bonusListId"`
				} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"shoulder"`
		Back struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem      int `json:"transmogItem"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				ItemAppearanceModID         int `json:"itemAppearanceModId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"back"`
		Chest struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem      int `json:"transmogItem"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				ItemAppearanceModID         int `json:"itemAppearanceModId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []struct {
					ID          int `json:"id"`
					Tier        int `json:"tier"`
					SpellID     int `json:"spellId"`
					BonusListID int `json:"bonusListId"`
				} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"chest"`
		Tabard struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem      int `json:"transmogItem"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats                []interface{} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []interface{} `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				ItemAppearanceModID         int `json:"itemAppearanceModId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"tabard"`
		Wrist struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				Gem0              int `json:"gem0"`
				TransmogItem      int `json:"transmogItem"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				ItemAppearanceModID         int `json:"itemAppearanceModId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"wrist"`
		Hands struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem      int `json:"transmogItem"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				ItemAppearanceModID         int `json:"itemAppearanceModId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"hands"`
		Waist struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				Gem0              int `json:"gem0"`
				TransmogItem      int `json:"transmogItem"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				ItemAppearanceModID         int `json:"itemAppearanceModId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"waist"`
		Legs struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem      int `json:"transmogItem"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				ItemAppearanceModID         int `json:"itemAppearanceModId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"legs"`
		Feet struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem      int `json:"transmogItem"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				ItemAppearanceModID         int `json:"itemAppearanceModId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"feet"`
		Finger1 struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				Enchant           int `json:"enchant"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				EnchantDisplayInfoID int `json:"enchantDisplayInfoId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"finger1"`
		Finger2 struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				Gem0              int `json:"gem0"`
				Enchant           int `json:"enchant"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				EnchantDisplayInfoID int `json:"enchantDisplayInfoId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"finger2"`
		Trinket1 struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"trinket1"`
		Trinket2 struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []interface{} `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"trinket2"`
		MainHand struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				Enchant           int `json:"enchant"`
				TransmogItem      int `json:"transmogItem"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor      int `json:"armor"`
			WeaponInfo struct {
				Damage struct {
					Min      int `json:"min"`
					Max      int `json:"max"`
					ExactMin int `json:"exactMin"`
					ExactMax int `json:"exactMax"`
				} `json:"damage"`
				WeaponSpeed float64 `json:"weaponSpeed"`
				Dps         float64 `json:"dps"`
			} `json:"weaponInfo"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				EnchantDisplayInfoID        int `json:"enchantDisplayInfoId"`
				ItemAppearanceModID         int `json:"itemAppearanceModId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"mainHand"`
		OffHand struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				Enchant           int `json:"enchant"`
				TransmogItem      int `json:"transmogItem"`
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor      int `json:"armor"`
			WeaponInfo struct {
				Damage struct {
					Min      int `json:"min"`
					Max      int `json:"max"`
					ExactMin int `json:"exactMin"`
					ExactMax int `json:"exactMax"`
				} `json:"damage"`
				WeaponSpeed float64 `json:"weaponSpeed"`
				Dps         float64 `json:"dps"`
			} `json:"weaponInfo"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				EnchantDisplayInfoID        int `json:"enchantDisplayInfoId"`
				ItemAppearanceModID         int `json:"itemAppearanceModId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
			AzeriteItem struct {
				AzeriteLevel               int `json:"azeriteLevel"`
				AzeriteExperience          int `json:"azeriteExperience"`
				AzeriteExperienceRemaining int `json:"azeriteExperienceRemaining"`
			} `json:"azeriteItem"`
			AzeriteEmpoweredItem struct {
				AzeritePowers []interface{} `json:"azeritePowers"`
			} `json:"azeriteEmpoweredItem"`
		} `json:"offHand"`
	} `json:"items"`
	TotalHonorableKills int `json:"totalHonorableKills"`
}
