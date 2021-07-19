

// 1:
   SetFontLoader sets a loader used to read font files (.json and .z) from an arbitrary source. 
   If a font loader has been specified, it is used to load the named font resources when AddFont() is called. 
   If this operation fails, an attempt is made to load the resources from the configured font directory


// 2: 

   AddUTF8Font imports a TrueType font with utf-8 symbols and makes it available.
   It is necessary to generate a font definition file first with the makefont
   utility. It is not necessary to call this function for the core PDF fonts
   (courier, helvetica, times, zapfdingbats).
   
   The JSON definition file (and the font file itself when embedding) must be
   present in the font directory. If it is not found, the error "Could not
   include font definition file" is set.
   
   family specifies the font family. The name can be chosen arbitrarily. If it
   is a standard family name, it will override the corresponding font. This
   string is used to subsequently set the font with the SetFont method.
   
   style specifies the font style. Acceptable values are (case insensitive) the
   empty string for regular style, "B" for bold, "I" for italic, or "BI" or
   "IB" for bold and italic combined.
   
   fileStr specifies the base name with ".json" extension of the font
   definition file to be added. The file will be loaded from the font directory
   specified in the call to New() or SetFontLocation().