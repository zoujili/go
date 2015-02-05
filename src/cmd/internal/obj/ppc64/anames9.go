package ppc64

/*
 * this is the ranlib header
 */
var Anames = []string{
	"XXX ",
	"CALL",
	"CHECKNIL",
	"DATA",
	"DUFFCOPY",
	"DUFFZERO",
	"END",
	"FUNCDATA",
	"GLOBL",
	"JMP",
	"NOP",
	"PCDATA",
	"RET",
	"TEXT",
	"TYPE",
	"UNDEF",
	"USEFIELD",
	"VARDEF",
	"VARKILL",
	"ADD ",
	"ADDCC",
	"ADDV",
	"ADDVCC",
	"ADDC",
	"ADDCCC",
	"ADDCV",
	"ADDCVCC",
	"ADDME",
	"ADDMECC",
	"ADDMEVCC",
	"ADDMEV",
	"ADDE",
	"ADDECC",
	"ADDEVCC",
	"ADDEV",
	"ADDZE",
	"ADDZECC",
	"ADDZEVCC",
	"ADDZEV",
	"AND",
	"ANDCC",
	"ANDN",
	"ANDNCC",
	"BC",
	"BCL",
	"BEQ",
	"BGE",
	"BGT",
	"BLE",
	"BLT",
	"BNE",
	"BVC",
	"BVS",
	"CMP",
	"CMPU",
	"CNTLZW",
	"CNTLZWCC",
	"CRAND",
	"CRANDN",
	"CREQV",
	"CRNAND",
	"CRNOR",
	"CROR",
	"CRORN",
	"CRXOR",
	"DIVW",
	"DIVWCC",
	"DIVWVCC",
	"DIVWV",
	"DIVWU",
	"DIVWUCC",
	"DIVWUVCC",
	"DIVWUV",
	"EQV",
	"EQVCC",
	"EXTSB",
	"EXTSBCC",
	"EXTSH",
	"EXTSHCC",
	"FABS",
	"FABSCC",
	"FADD",
	"FADDCC",
	"FADDS",
	"FADDSCC",
	"FCMPO",
	"FCMPU",
	"FCTIW",
	"FCTIWCC",
	"FCTIWZ",
	"FCTIWZCC",
	"FDIV",
	"FDIVCC",
	"FDIVS",
	"FDIVSCC",
	"FMADD",
	"FMADDCC",
	"FMADDS",
	"FMADDSCC",
	"FMOVD",
	"FMOVDCC",
	"FMOVDU",
	"FMOVS",
	"FMOVSU",
	"FMSUB",
	"FMSUBCC",
	"FMSUBS",
	"FMSUBSCC",
	"FMUL",
	"FMULCC",
	"FMULS",
	"FMULSCC",
	"FNABS",
	"FNABSCC",
	"FNEG",
	"FNEGCC",
	"FNMADD",
	"FNMADDCC",
	"FNMADDS",
	"FNMADDSCC",
	"FNMSUB",
	"FNMSUBCC",
	"FNMSUBS",
	"FNMSUBSCC",
	"FRSP",
	"FRSPCC",
	"FSUB",
	"FSUBCC",
	"FSUBS",
	"FSUBSCC",
	"MOVMW",
	"LSW",
	"LWAR",
	"MOVWBR",
	"MOVB",
	"MOVBU",
	"MOVBZ",
	"MOVBZU",
	"MOVH",
	"MOVHBR",
	"MOVHU",
	"MOVHZ",
	"MOVHZU",
	"MOVW",
	"MOVWU",
	"MOVFL",
	"MOVCRFS",
	"MTFSB0",
	"MTFSB0CC",
	"MTFSB1",
	"MTFSB1CC",
	"MULHW",
	"MULHWCC",
	"MULHWU",
	"MULHWUCC",
	"MULLW",
	"MULLWCC",
	"MULLWVCC",
	"MULLWV",
	"NAND",
	"NANDCC",
	"NEG",
	"NEGCC",
	"NEGVCC",
	"NEGV",
	"NOR",
	"NORCC",
	"OR",
	"ORCC",
	"ORN",
	"ORNCC",
	"REM",
	"REMCC",
	"REMV",
	"REMVCC",
	"REMU",
	"REMUCC",
	"REMUV",
	"REMUVCC",
	"RFI",
	"RLWMI",
	"RLWMICC",
	"RLWNM",
	"RLWNMCC",
	"SLW",
	"SLWCC",
	"SRW",
	"SRAW",
	"SRAWCC",
	"SRWCC",
	"STSW",
	"STWCCC",
	"SUB",
	"SUBCC",
	"SUBVCC",
	"SUBC",
	"SUBCCC",
	"SUBCV",
	"SUBCVCC",
	"SUBME",
	"SUBMECC",
	"SUBMEVCC",
	"SUBMEV",
	"SUBV",
	"SUBE",
	"SUBECC",
	"SUBEV",
	"SUBEVCC",
	"SUBZE",
	"SUBZECC",
	"SUBZEVCC",
	"SUBZEV",
	"SYNC",
	"XOR",
	"XORCC",
	"DCBF",
	"DCBI",
	"DCBST",
	"DCBT",
	"DCBTST",
	"DCBZ",
	"ECIWX",
	"ECOWX",
	"EIEIO",
	"ICBI",
	"ISYNC",
	"PTESYNC",
	"TLBIE",
	"TLBIEL",
	"TLBSYNC",
	"TW",
	"SYSCALL",
	"WORD",
	"RFCI",
	"FRES",
	"FRESCC",
	"FRSQRTE",
	"FRSQRTECC",
	"FSEL",
	"FSELCC",
	"FSQRT",
	"FSQRTCC",
	"FSQRTS",
	"FSQRTSCC",
	"CNTLZD",
	"CNTLZDCC",
	"CMPW",
	"CMPWU",
	"DIVD",
	"DIVDCC",
	"DIVDVCC",
	"DIVDV",
	"DIVDU",
	"DIVDUCC",
	"DIVDUVCC",
	"DIVDUV",
	"EXTSW",
	"EXTSWCC",
	"FCFID",
	"FCFIDCC",
	"FCTID",
	"FCTIDCC",
	"FCTIDZ",
	"FCTIDZCC",
	"LDAR",
	"MOVD",
	"MOVDU",
	"MOVWZ",
	"MOVWZU",
	"MULHD",
	"MULHDCC",
	"MULHDU",
	"MULHDUCC",
	"MULLD",
	"MULLDCC",
	"MULLDVCC",
	"MULLDV",
	"RFID",
	"RLDMI",
	"RLDMICC",
	"RLDC",
	"RLDCCC",
	"RLDCR",
	"RLDCRCC",
	"RLDCL",
	"RLDCLCC",
	"SLBIA",
	"SLBIE",
	"SLBMFEE",
	"SLBMFEV",
	"SLBMTE",
	"SLD",
	"SLDCC",
	"SRD",
	"SRAD",
	"SRADCC",
	"SRDCC",
	"STDCCC",
	"TD",
	"DWORD",
	"REMD",
	"REMDCC",
	"REMDV",
	"REMDVCC",
	"REMDU",
	"REMDUCC",
	"REMDUV",
	"REMDUVCC",
	"HRFID",
	"LAST",
}

var cnames9 = []string{
	"NONE",
	"REG",
	"FREG",
	"CREG",
	"SPR",
	"ZCON",
	"SCON",
	"UCON",
	"ADDCON",
	"ANDCON",
	"LCON",
	"DCON",
	"SACON",
	"SECON",
	"LACON",
	"LECON",
	"DACON",
	"SBRA",
	"LBRA",
	"SAUTO",
	"LAUTO",
	"SEXT",
	"LEXT",
	"ZOREG",
	"SOREG",
	"LOREG",
	"FPSCR",
	"MSR",
	"XER",
	"LR",
	"CTR",
	"ANY",
	"GOK",
	"ADDR",
	"TEXTSIZE",
	"NCLASS",
}
