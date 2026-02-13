# Neovim was the catalyst
> 02-12-26

It’s been years since the last time I wrote something and genuinely felt the need to immortalize the experience I had. Not to exaggerate, but neovim has truly been life changing.

Back in the day, I used to get overwhelmed with GUI clients i.e., for databases, so I tried and made the switch to terminal. Long story short, the same transition eventually happened with Git, Maven, package managers, and more. 

Looking back on nearly a decade of professional experience, it’s probably been 8 years since I last installed a GUI client for those tools. And now, it happened again, this time with the IDEs I’ve grown so used to.

Story time. 

I’ve been using JetBrains (IntelliJ) for most of my career working with Spring projects. It’s just that good TBH, and I believe you can’t go wrong with it. Oddly enough (or maybe as common as it can be), I default to VS code when working with languages other than Java. 

This wasn't my first time trying neovim. In fact, I once gave up on it and ended up installing vim motions in my go-to IDEs. 

Fast forward to January 2026. I decided to give it another shot and really put it to the test by actually using it at work. It was still a stock neovim setup at the time. (I know. Crazy.)

I wish I could say I delivered the requirement faster than ever, but far from it, SKILL ISSUE hit me hard. I had to quickly abort mission, switch back to IntelliJ, and prioritize the product requirements. 

While it obviously didn’t go as planned (yet), I realized something very important, and that is how reliant I had become on IntelliJ, and how much I had been taking IDE features for granted.

A huge part of coding in an IDE had become muscle memory. I guess getting so caught up with work, life, and everything in between over the years made me so numb. Thankfully, that single experience alone made me appreciate those little things again. It rekindled a tiny ember in me, to setup an editor I can consider my own. 

So while coding on an IDE, I watched myself code as if I were on a third person POV, and without being on auto pilot, I noticed a few patterns in how I worked, and so I set those up in neovim (even learned a bit of lua in the process).

The next day, I put it to the test again. 

I thought it was gonna be different, but no. I was still missing a lot and it turns out setting up neovim for Java isn’t the most straightforward process. I cycled through this process every day for two weeks and ended up with the ff setup:
- Built my own basic boilerplate generator for creating classes, enums, interfaces, and records in .java files
- **neo-tree** for file trees (immediately questioned myself, more on that later)
- **telescope** for fuzzy search
- **JDTLS** setup and configured LSP keymaps (this took longer than I'd want to.)

At that point, I could at least get by, but it still wasn’t quite there. So I kept going.

Over the next two weeks:
- Moved from neo-tree to nvim-tree
- Then had an internal debate about file trees altogether
Part of me tried to rationalize the need for them. Another part saw how deep nesting is often just noise. Initially, I told myself I needed file trees to build a better mental model of the project structure. But to be sure, I observed myself again.

Oddly enough, I noticed I had a tendency to open and stare at the file tree whenever I was thinking through a problem.

There’s nothing inherently wrong with that. In fact, I realized it subconsciously helped me build a mental map of related components. But I wanted a more intentional way of coding—focusing on the relevant code that needed solving, and using a flatter file explorer only when necessary.

So I moved from Nvim-tree to oil.nvim.

Then came another internal debate. JDTLS’s code action for moving a type to another package felt counterproductive. I came up with a hacky idea I was willing to settle for—minimal repackaging support. So I forked oil.nvim, implemented an adjustment that exports the file rename event, and leveraged it to update the package name (and dependencies when needed) whenever a type is moved.

At this point, with no file tree and no tabs, the need to cycle through specific files surfaced. That’s when I looked into Harpoon.

After that, I started experimenting with modern terminal emulators. I switched from iTerm2 to Ghostty. Then I automated app launching with specific window dimensions using hotkeys, Raycast and AppleScript did the trick. Since the terminal is one of those apps, I included a tmux attach-session command in the AppleScript, effectively making Ghostty a session-aware terminal-based environment.

Then I discovered Git Worktrees, which further improved my quality of life.

I’m not a fan of multitasking, and never will be. I firmly believe that, at our core, we do things one at a time. Even with this setup in place, that remains true for me.

But if there’s one thing this setup truly nails, it’s cutting down a huge portion of the mechanical overhead of context switching.

Now, I’ve fully transitioned to Neovim. I uninstalled VS Code and only removed IntelliJ from the Dock. I’ve made peace with the fact that if I need to do massive refactors, like moving multiple modules across packages, IntelliJ is still the better tool for the job.

And that’s okay.
