package ei

var _familySortKeys = map[ArtifactSpec_Name]uint32{
	ArtifactSpec_PUZZLE_CUBE:          1,
	ArtifactSpec_LUNAR_TOTEM:          2,
	ArtifactSpec_DEMETERS_NECKLACE:    3,
	ArtifactSpec_VIAL_MARTIAN_DUST:    4,
	ArtifactSpec_AURELIAN_BROOCH:      5,
	ArtifactSpec_TUNGSTEN_ANKH:        6,
	ArtifactSpec_ORNATE_GUSSET:        7,
	ArtifactSpec_NEODYMIUM_MEDALLION:  8,
	ArtifactSpec_MERCURYS_LENS:        9,
	ArtifactSpec_BEAK_OF_MIDAS:        10,
	ArtifactSpec_CARVED_RAINSTICK:     11,
	ArtifactSpec_INTERSTELLAR_COMPASS: 12,
	ArtifactSpec_THE_CHALICE:          13,
	ArtifactSpec_PHOENIX_FEATHER:      14,
	ArtifactSpec_QUANTUM_METRONOME:    15,
	ArtifactSpec_DILITHIUM_MONOCLE:    16,
	ArtifactSpec_TITANIUM_ACTUATOR:    17,
	ArtifactSpec_SHIP_IN_A_BOTTLE:     18,
	ArtifactSpec_TACHYON_DEFLECTOR:    19,
	ArtifactSpec_BOOK_OF_BASAN:        20,
	ArtifactSpec_LIGHT_OF_EGGENDIL:    21,

	ArtifactSpec_LUNAR_STONE:     101,
	ArtifactSpec_SHELL_STONE:     102,
	ArtifactSpec_TACHYON_STONE:   103,
	ArtifactSpec_TERRA_STONE:     104,
	ArtifactSpec_SOUL_STONE:      105,
	ArtifactSpec_DILITHIUM_STONE: 106,
	ArtifactSpec_QUANTUM_STONE:   107,
	ArtifactSpec_LIFE_STONE:      108,
	ArtifactSpec_CLARITY_STONE:   109,
	ArtifactSpec_PROPHECY_STONE:  110,

	ArtifactSpec_GOLD_METEORITE: 201,
	ArtifactSpec_TAU_CETI_GEODE: 202,
	ArtifactSpec_SOLAR_TITANIUM: 203,
}

func FamilySortKey(family ArtifactSpec_Name) uint32 {
	return _familySortKeys[family]
}
