# Neovim was the catalyst
> 02-12-26

It’s been years since the last time I wrote something and genuinely felt the need to immortalize the experience I had. Not to exaggerate, but neovim has truly been life changing.

Back in the day, I used to get overwhelmed with GUI clients i.e., for databases, so I tried and made the switch to terminal. Long story short, the same transition happened with Git, Maven, package managers, and more. 

Looking back on nearly a decade of professional experience, it’s probably been 8 years since I last installed a GUI client for those tools. Now, it happened again, this time with the IDEs I’ve grown so used to.

Story time. I’ve been using JetBrains (IntelliJ) for the longest time through out my career working with Spring projects. It’s just that good TBH, and I believe you can’t go wrong with it. Oddly enough (or maybe as common as it can be), I default to using VS code when working with languages other than Java. It’s not my first time trying neovim. In fact, I once gave up on it and ended up installing vim motions in my go-to IDEs. 

Fast forward to January 2026. I decided to give it another shot and really put it to the test by actually using it at work (it was a still a stock neovim setup at the time. I know. Crazy.)

I wish I could say I delivered the requirement faster than ever, but far from it, SKILL ISSUE hit me hard. I had to quickly abort the mission and switch back to IntelliJ, and prioritize the product requirements. 

While it obviously didn’t go as planned (yet), I realized something very important, and that is how reliant I had become on IntelliJ, and how much I had been taking IDE features for granted.

A huge part of coding in an IDE had become muscle memory for me and I guess getting so caught up with work, life, and any shenanigans in between throughout the years made me so numb. Thankfully, that single experience alone made me appreciate those little things again. It rekindled a tiny ember in me, to setup an editor I can consider my own. 

So with me on an IDE, I watched myself code as if I were on a third person POV, and without being on auto pilot, I noticed a few things and had set those up in neovim (even learned a bit of lua in the process).

The next day, I tried putting it to the test again. I thought it was gonna be different this time, but no, it turns out I'm still missing a lot and setting up neovim for Java isn’t the most straightforward process. I cycled through this experience every day for two weeks and ended up with the ff setup:
- made my own basic boiler plate generator when creating classes/enums/interfaces/records for .java files
- neo-tree for file trees (immediately questioned myself, more on that later)
- telescope for fuzzy search
- jdtls setup and configure lsp keymaps (this took longer than I'd want to.)

I find myself being able to at least get by at this point, but it's definitely wasn't quite there yet. So I kept going, and for the next two weeks:
- moved from neo-tree to nvim-tree
- then had an internal debate with myself regarding file trees, where a part of me tries to rationalize the need for it and the other part of me sees how the deep nesting is basically just noise most of the time. At first, I kept telling myself that I'm using file trees to have a better model of the project structure, but just to be sure, I tried observing myself again. Oddly enough, I found out that I had this tendency to open and stare at the file tree whenver thinking about a problem. While there's nothing inherently wrong with it, infact I realized that it was through that I subconsciously get to build the mental model of related components, I thought to myself that I wanted a more intentional way of coding i.e., focus on the relevant code that needs solving, use a flat file explorer as needed. And so I moved from nvim-tree to oil.nvim.
- had another internal debate with myself because of how counter productive jdtls' code action to move a type to another package. Thought of a hacky idea I'm willing to settle on for minimal repackaging. So I forked oil.nvim, implemented an adjustment that exports the file rename event, leveraged the rename event to update the package name when a type is moved and its dependencies as needed.
- At this point, since I had no access to a file tree or to tabs, the need cycle through specific files surfaces, and that's when I thought of checking harpoon.

After those, I thought of trying it out modern terminal emulators and so I made the switch from iTerm2 to Ghostty. Then thought of finally automating launching of apps with specified dimensions with the use hot keys. For that, I used raycast and applescript. Since the terminal is one of these apps, I included tmux attach session command to its applescript making Ghostty a session-aware terminal-based environment. I then came across Git Worktrees and it also has improved my quality of life even more.    

Not a fan of multi tasking, and will never be, as I firmly believe that at our core as human beings, we're simply doing things one at a time. And even With this setup in place, that remains true for me. But if there's one thing that this setup truly nails, it's cutting down a huge portion of mechanical parts of context switching.

Now, I have fully transitioned to neovim. Uninstalled VS code but only removed IntelliJ from the Dock. I've made peace that if I need to do huge refactor i.e., moving a bunch of modules to different packages, then I'd better use IntelliJ for that. 
