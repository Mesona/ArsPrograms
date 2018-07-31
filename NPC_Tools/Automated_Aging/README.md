This program is built to be used in conjunction with the spreadsheet that can be copied here:
https://drive.google.com/file/d/1C5L94AwmdRPBpqQsx74oC3ikz9HO6Cf1/view?usp=sharing

Simply download the "Aging Chart Template.csv" and the main.go, put both files in the same folder on your machine, then run "go run main.go."  If you do not rename "Aging Chart Template.csv" everything should work out of the box, but if you do rename the file, you'll need to edit the variable "filename" at line 19 of main.go.

Some features were left in this NPC tool because I got slightly carried away and started including elements related to PC age advancement, but the core system works.  My group doesn't track the individual characteristics of our townsfolk, so this program only tracks decrepitude.  At some point I'll possibly work on creating a PC aging tool that's more robust, but that's a very low priority.

You will need to manually set the current date of your game, as well as put in NPC details when you first recruit them, but after that the aging will be handled automatically.  NPCs are not removed when they die, they will need to be removed manually, but the "Crisis" column will show a helpful reminder if any NPC's decrepitude score hits 5.

You will also need to set whatever your covenant bonuses to aging are as the "Living Modifiers."  These are things such as Wealthy (-2 to aging rolls) or whatever your covenant has, and which affects everyone uniformly.

Thirdly, if your NPC has any of the virtues or flaws that affect aging from the core book (Faerie Blood, Strong Faerie Blood, Unaging, and Quick Aging), those can be denoted in the "Virtues/Flaws" field with FB SFB UA or QA.  The field can handle multiple, so if your NPC has both Faerie Blood and Quick Aging, you can out "FBQA" in that field and the program will factor those both in.  And yes, I know NPCs can't normally have Strong Faerie Blood, but I added it already so it's staying in.

Finally, if your NPC has any personal aging modifiers, either positive or negative, that are not covered by feats, you can add those in the "Personal Modifiers" tab.
