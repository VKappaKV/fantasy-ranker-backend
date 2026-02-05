package domain

type Champion string

const (
	ChampionAatrox       Champion = "Aatrox"
	ChampionAhri         Champion = "Ahri"
	ChampionAkali        Champion = "Akali"
	ChampionAkshan       Champion = "Akshan"
	ChampionAlistar      Champion = "Alistar"
	ChampionAmbessa      Champion = "Ambessa"
	ChampionAmumu        Champion = "Amumu"
	ChampionAnivia       Champion = "Anivia"
	ChampionAnnie        Champion = "Annie"
	ChampionAphelios     Champion = "Aphelios"
	ChampionAshe         Champion = "Ashe"
	ChampionAurelionSol  Champion = "AurelionSol"
	ChampionAurora       Champion = "Aurora"
	ChampionAzir         Champion = "Azir"
	ChampionBard         Champion = "Bard"
	ChampionBelveth      Champion = "Belveth"
	ChampionBlitzcrank   Champion = "Blitzcrank"
	ChampionBrand        Champion = "Brand"
	ChampionBraum        Champion = "Braum"
	ChampionBriar        Champion = "Briar"
	ChampionCaitlyn      Champion = "Caitlyn"
	ChampionCamille      Champion = "Camille"
	ChampionCassiopeia   Champion = "Cassiopeia"
	ChampionChogath      Champion = "Chogath"
	ChampionCorki        Champion = "Corki"
	ChampionDarius       Champion = "Darius"
	ChampionDiana        Champion = "Diana"
	ChampionDraven       Champion = "Draven"
	ChampionDrMundo      Champion = "DrMundo"
	ChampionEkko         Champion = "Ekko"
	ChampionElise        Champion = "Elise"
	ChampionEvelynn      Champion = "Evelynn"
	ChampionEzreal       Champion = "Ezreal"
	ChampionFiddlesticks Champion = "Fiddlesticks"
	ChampionFiora        Champion = "Fiora"
	ChampionFizz         Champion = "Fizz"
	ChampionGalio        Champion = "Galio"
	ChampionGangplank    Champion = "Gangplank"
	ChampionGaren        Champion = "Garen"
	ChampionGnar         Champion = "Gnar"
	ChampionGragas       Champion = "Gragas"
	ChampionGraves       Champion = "Graves"
	ChampionGwen         Champion = "Gwen"
	ChampionHecarim      Champion = "Hecarim"
	ChampionHeimerdinger Champion = "Heimerdinger"
	ChampionHwei         Champion = "Hwei"
	ChampionIllaoi       Champion = "Illaoi"
	ChampionIrelia       Champion = "Irelia"
	ChampionIvern        Champion = "Ivern"
	ChampionJanna        Champion = "Janna"
	ChampionJarvanIV     Champion = "JarvanIV"
	ChampionJax          Champion = "Jax"
	ChampionJayce        Champion = "Jayce"
	ChampionJhin         Champion = "Jhin"
	ChampionJinx         Champion = "Jinx"
	ChampionKaisa        Champion = "Kaisa"
	ChampionKalista      Champion = "Kalista"
	ChampionKarma        Champion = "Karma"
	ChampionKarthus      Champion = "Karthus"
	ChampionKassadin     Champion = "Kassadin"
	ChampionKatarina     Champion = "Katarina"
	ChampionKayle        Champion = "Kayle"
	ChampionKayn         Champion = "Kayn"
	ChampionKennen       Champion = "Kennen"
	ChampionKhazix       Champion = "Khazix"
	ChampionKindred      Champion = "Kindred"
	ChampionKled         Champion = "Kled"
	ChampionKogMaw       Champion = "KogMaw"
	ChampionKSante       Champion = "KSante"
	ChampionLeblanc      Champion = "Leblanc"
	ChampionLeeSin       Champion = "LeeSin"
	ChampionLeona        Champion = "Leona"
	ChampionLillia       Champion = "Lillia"
	ChampionLissandra    Champion = "Lissandra"
	ChampionLucian       Champion = "Lucian"
	ChampionLulu         Champion = "Lulu"
	ChampionLux          Champion = "Lux"
	ChampionMalphite     Champion = "Malphite"
	ChampionMalzahar     Champion = "Malzahar"
	ChampionMaokai       Champion = "Maokai"
	ChampionMasterYi     Champion = "MasterYi"
	ChampionMel          Champion = "Mel"
	ChampionMilio        Champion = "Milio"
	ChampionMissFortune  Champion = "MissFortune"
	ChampionMonkeyKing   Champion = "MonkeyKing"
	ChampionMordekaiser  Champion = "Mordekaiser"
	ChampionMorgana      Champion = "Morgana"
	ChampionNaafiri      Champion = "Naafiri"
	ChampionNami         Champion = "Nami"
	ChampionNasus        Champion = "Nasus"
	ChampionNautilus     Champion = "Nautilus"
	ChampionNeeko        Champion = "Neeko"
	ChampionNidalee      Champion = "Nidalee"
	ChampionNilah        Champion = "Nilah"
	ChampionNocturne     Champion = "Nocturne"
	ChampionNunu         Champion = "Nunu"
	ChampionOlaf         Champion = "Olaf"
	ChampionOrianna      Champion = "Orianna"
	ChampionOrnn         Champion = "Ornn"
	ChampionPantheon     Champion = "Pantheon"
	ChampionPoppy        Champion = "Poppy"
	ChampionPyke         Champion = "Pyke"
	ChampionQiyana       Champion = "Qiyana"
	ChampionQuinn        Champion = "Quinn"
	ChampionRakan        Champion = "Rakan"
	ChampionRammus       Champion = "Rammus"
	ChampionRekSai       Champion = "RekSai"
	ChampionRell         Champion = "Rell"
	ChampionRenata       Champion = "Renata"
	ChampionRenekton     Champion = "Renekton"
	ChampionRengar       Champion = "Rengar"
	ChampionRiven        Champion = "Riven"
	ChampionRumble       Champion = "Rumble"
	ChampionRyze         Champion = "Ryze"
	ChampionSamira       Champion = "Samira"
	ChampionSejuani      Champion = "Sejuani"
	ChampionSenna        Champion = "Senna"
	ChampionSeraphine    Champion = "Seraphine"
	ChampionSett         Champion = "Sett"
	ChampionShaco        Champion = "Shaco"
	ChampionShen         Champion = "Shen"
	ChampionShyvana      Champion = "Shyvana"
	ChampionSinged       Champion = "Singed"
	ChampionSion         Champion = "Sion"
	ChampionSivir        Champion = "Sivir"
	ChampionSkarner      Champion = "Skarner"
	ChampionSmolder      Champion = "Smolder"
	ChampionSona         Champion = "Sona"
	ChampionSoraka       Champion = "Soraka"
	ChampionSwain        Champion = "Swain"
	ChampionSylas        Champion = "Sylas"
	ChampionSyndra       Champion = "Syndra"
	ChampionTahmKench    Champion = "TahmKench"
	ChampionTaliyah      Champion = "Taliyah"
	ChampionTalon        Champion = "Talon"
	ChampionTaric        Champion = "Taric"
	ChampionTeemo        Champion = "Teemo"
	ChampionThresh       Champion = "Thresh"
	ChampionTristana     Champion = "Tristana"
	ChampionTrundle      Champion = "Trundle"
	ChampionTryndamere   Champion = "Tryndamere"
	ChampionTwistedFate  Champion = "TwistedFate"
	ChampionTwitch       Champion = "Twitch"
	ChampionUdyr         Champion = "Udyr"
	ChampionUrgot        Champion = "Urgot"
	ChampionVarus        Champion = "Varus"
	ChampionVayne        Champion = "Vayne"
	ChampionVeigar       Champion = "Veigar"
	ChampionVelkoz       Champion = "Velkoz"
	ChampionVex          Champion = "Vex"
	ChampionVi           Champion = "Vi"
	ChampionViego        Champion = "Viego"
	ChampionViktor       Champion = "Viktor"
	ChampionVladimir     Champion = "Vladimir"
	ChampionVolibear     Champion = "Volibear"
	ChampionWarwick      Champion = "Warwick"
	ChampionXayah        Champion = "Xayah"
	ChampionXerath       Champion = "Xerath"
	ChampionXinZhao      Champion = "XinZhao"
	ChampionYasuo        Champion = "Yasuo"
	ChampionYone         Champion = "Yone"
	ChampionYorick       Champion = "Yorick"
	ChampionYunara       Champion = "Yunara"
	ChampionYuumi        Champion = "Yuumi"
	ChampionZaahen       Champion = "Zaahen"
	ChampionZac          Champion = "Zac"
	ChampionZed          Champion = "Zed"
	ChampionZeri         Champion = "Zeri"
	ChampionZiggs        Champion = "Ziggs"
	ChampionZilean       Champion = "Zilean"
	ChampionZoe          Champion = "Zoe"
	ChampionZyra         Champion = "Zyra"
)

// AllChampions contains every champion ID as defined by Data Dragon champion.json.
var AllChampions = []Champion{
	ChampionAatrox,
	ChampionAhri,
	ChampionAkali,
	ChampionAkshan,
	ChampionAlistar,
	ChampionAmbessa,
	ChampionAmumu,
	ChampionAnivia,
	ChampionAnnie,
	ChampionAphelios,
	ChampionAshe,
	ChampionAurelionSol,
	ChampionAurora,
	ChampionAzir,
	ChampionBard,
	ChampionBelveth,
	ChampionBlitzcrank,
	ChampionBrand,
	ChampionBraum,
	ChampionBriar,
	ChampionCaitlyn,
	ChampionCamille,
	ChampionCassiopeia,
	ChampionChogath,
	ChampionCorki,
	ChampionDarius,
	ChampionDiana,
	ChampionDraven,
	ChampionDrMundo,
	ChampionEkko,
	ChampionElise,
	ChampionEvelynn,
	ChampionEzreal,
	ChampionFiddlesticks,
	ChampionFiora,
	ChampionFizz,
	ChampionGalio,
	ChampionGangplank,
	ChampionGaren,
	ChampionGnar,
	ChampionGragas,
	ChampionGraves,
	ChampionGwen,
	ChampionHecarim,
	ChampionHeimerdinger,
	ChampionHwei,
	ChampionIllaoi,
	ChampionIrelia,
	ChampionIvern,
	ChampionJanna,
	ChampionJarvanIV,
	ChampionJax,
	ChampionJayce,
	ChampionJhin,
	ChampionJinx,
	ChampionKaisa,
	ChampionKalista,
	ChampionKarma,
	ChampionKarthus,
	ChampionKassadin,
	ChampionKatarina,
	ChampionKayle,
	ChampionKayn,
	ChampionKennen,
	ChampionKhazix,
	ChampionKindred,
	ChampionKled,
	ChampionKogMaw,
	ChampionKSante,
	ChampionLeblanc,
	ChampionLeeSin,
	ChampionLeona,
	ChampionLillia,
	ChampionLissandra,
	ChampionLucian,
	ChampionLulu,
	ChampionLux,
	ChampionMalphite,
	ChampionMalzahar,
	ChampionMaokai,
	ChampionMasterYi,
	ChampionMel,
	ChampionMilio,
	ChampionMissFortune,
	ChampionMonkeyKing,
	ChampionMordekaiser,
	ChampionMorgana,
	ChampionNaafiri,
	ChampionNami,
	ChampionNasus,
	ChampionNautilus,
	ChampionNeeko,
	ChampionNidalee,
	ChampionNilah,
	ChampionNocturne,
	ChampionNunu,
	ChampionOlaf,
	ChampionOrianna,
	ChampionOrnn,
	ChampionPantheon,
	ChampionPoppy,
	ChampionPyke,
	ChampionQiyana,
	ChampionQuinn,
	ChampionRakan,
	ChampionRammus,
	ChampionRekSai,
	ChampionRell,
	ChampionRenata,
	ChampionRenekton,
	ChampionRengar,
	ChampionRiven,
	ChampionRumble,
	ChampionRyze,
	ChampionSamira,
	ChampionSejuani,
	ChampionSenna,
	ChampionSeraphine,
	ChampionSett,
	ChampionShaco,
	ChampionShen,
	ChampionShyvana,
	ChampionSinged,
	ChampionSion,
	ChampionSivir,
	ChampionSkarner,
	ChampionSmolder,
	ChampionSona,
	ChampionSoraka,
	ChampionSwain,
	ChampionSylas,
	ChampionSyndra,
	ChampionTahmKench,
	ChampionTaliyah,
	ChampionTalon,
	ChampionTaric,
	ChampionTeemo,
	ChampionThresh,
	ChampionTristana,
	ChampionTrundle,
	ChampionTryndamere,
	ChampionTwistedFate,
	ChampionTwitch,
	ChampionUdyr,
	ChampionUrgot,
	ChampionVarus,
	ChampionVayne,
	ChampionVeigar,
	ChampionVelkoz,
	ChampionVex,
	ChampionVi,
	ChampionViego,
	ChampionViktor,
	ChampionVladimir,
	ChampionVolibear,
	ChampionWarwick,
	ChampionXayah,
	ChampionXerath,
	ChampionXinZhao,
	ChampionYasuo,
	ChampionYone,
	ChampionYorick,
	ChampionYunara,
	ChampionYuumi,
	ChampionZaahen,
	ChampionZac,
	ChampionZed,
	ChampionZeri,
	ChampionZiggs,
	ChampionZilean,
	ChampionZoe,
	ChampionZyra,
}

// todo: add all champions