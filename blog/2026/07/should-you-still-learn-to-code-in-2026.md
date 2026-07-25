---
title: Should You Still Learn to Code in 2026?
subtitle: The answer isn’t as obvious as I used to believe.
tags:
- Software Engineering
- Coding
- Python
- Programming
- Computer Science
published: '2026-02-23'
updated: '2026-05-29'
free: false
freedium_url: https://freedium-mirror.cfd/https://medium.com/data-science-collective/should-you-still-learn-to-code-in-2026-034685e17707
source_url: https://medium.com/data-science-collective/should-you-still-learn-to-code-in-2026-034685e17707
---

# Should You Still Learn to Code in 2026?

*The answer isn’t as obvious as I used to believe.*

*Published Feb 23, 2026 · Updated May 29, 2026 · Free: No*

<mark class="bg-emerald-300 dark:bg-emerald-700 dark:text-white">I'm a Senior Applied Scientist at Amazon</mark> where I build production machine learning systems. And I don't write code anymore.

AI writes virtually every line I commit.

So here's the uncomfortable question I keep getting, and honestly, the question I've been asking myself: if someone like me doesn't actually write code anymore… is it too late for you to learn?

Last year if you'd asked me I would have been totally confident that AI coding tools just aren't there yet, and I was skeptical if they'd ever be. Today, I think people who are skeptical about AI being able to code at a professional level just aren't using the tools right. In the right hands, AI assistants like Claude Code are WAY past just being auto-complete and now able to do complex, multi-step workflows that literally span days.

But this is just my experience. So I dug into the data.

And at first? It looks pretty bleak.

### What the Data Say

Let's start with the job market. The Bureau of Labor Statistics shows that "computer programmer" roles dropped about 27% in just two years, and they project another 6% decline through 2034. Those jobs straight up aren't coming back.

Not to mention layoffs and the correction of over-hiring during Covid. As of mid-2025, tech postings on Indeed are 36% below the pre-pandemic baseline.

All this while AI use is growing. Stack Overflow's 2025 developer survey found 84% of developers are using or planning to use AI tools. And brand new data from the Pragmatic Summit — a gathering of 500 top engineers — puts it even higher. 93% of devs are now using AI tools, saving an average of four hours a week. AI-authored code jumped from 22% in Q3 2025 to 27% by February 2026. That's a massive shift in just one quarter. AI can create entire applications from spec to testing to deployment.

So if you're watching this thinking, "Why would I spend years learning something AI can already do?," that's a fair reaction.

But the data tell a more complicated story than the headlines suggest.

Because here's what those scary numbers leave out.

"Computer programmer" roles are dying, but "software developer" roles are down 0.3%. That's basically flat. And the Bureau of Labor Statistics actually projects 15% growth for software developers 2034. That's five times faster than the average for all occupations.

So what's the difference between those two categories?

"Programmer" was historically about translating specs into syntax. They would take requirements and convert them into working code. That's the part AI is good at — and honestly, it was heading toward automation long before ChatGPT. AI just accelerated it.

"Developer" and "engineer" roles involve design decisions, reliability, trade-off analysis, cross-functional communication, and incident response. All of the things that require judgment.

The work that's disappearing was always going to disappear. The work that's staying requires a human brain (for now at least, which we'll get to).

And here's something else: while overall tech hiring is down, AI-related demand is moving in the opposite direction. Axios reported that mentions of AI skills in job postings rose 16% in just three months, even as overall tech hiring was down 27%. What we're seeing is more of a market shift than anything else.

Now, remember that stat about 84% of developers using AI? There's a follow-up that's really important.

Stack Overflow found that 46% of developers actively distrust AI-generated code, up from 31% the year before. Only 3% say they "highly trust" it. Two-thirds of developers say AI gives answers that are "almost right, but not quite" — which makes debugging _more_ time-consuming, not less. It's creating code that looks correct, but isn't.

I'm one of the 84% of developers using AI, but I'm also part of the 46% who actively distrust AI-generated code.

### How Jobs Have Changed

So let me show you what my job actually looks like now that I don't write code myself anymore.

Think about software work in three phases:

1. **Before code**: What are we building and why? What are the constraints — things like latency, cost, and privacy? What could go wrong? Who are the stakeholders, and why do they care about this? What are the politics and personalities between teams that determine what gets built?
2. **During code**: Writing the actual functions, modules, and tests.
3. **After code**: This is everything from deployment to monitoring to compliance to incident response and communicating everything to stakeholders. All the stuff that is required for production systems and for decision-making.

AI compressed the _during_ phase, but it didn't magically delete the before and after. It actually made them MORE of a focus than they were previously.

Now what a project looks like for me may be a couple of weeks of coordinating with stakeholders, gathering requirements, and writing really detailed specs. Then a day or two of working with an AI coding assistant to actually build the project. Then potentially several more weeks of testing, evaluating, and making sure I'm confident in what I'm shipping.

That first part is really important. You have to have a clear idea of what you're building for the AI to be successful. I honestly think this explains the remaining AI skeptics.

Used correctly, with these tools you can make incredibly fast progress. AI gets you 80% of the way there in record time. But that last 20% — building the right things and making it production-safe — is where the actual hard work has always been.

And if you don't understand systems deeply enough to evaluate that last 20%, you're shipping code you can't vouch for.

Because here's what doesn't change regardless of how good AI gets: when something breaks in production — when there's a security breach, a compliance violation, or an outage that costs the company boatloads of money — someone is accountable.

AI doesn't get paged at 3am. You do.

AI doesn't get called into the incident review. You do.

AI doesn't explain to leadership why customer data was exposed. You do.

### Looking To The Future

So even in the most optimistic AI future, the question isn't "will humans be involved?" It's "what will humans need to know to be involved effectively?"

And the answer is: you need to understand systems. Which means you need to understand code.

You can't audit AI-generated code if you don't know what "correct" looks like. You can't debug a production incident if you can't read logs and stack traces. You can't make good architectural decisions if you don't understand databases, networking, concurrency, and failure modes.

It's not about the typing. It's about the understanding.

As Dave Farley from Modern Software Engineering put it, AI code assistance acts as a kind of amplifier. If you're already doing the right things, AI will amplify those things. If you're already doing the wrong things, AI will help you to dig a deeper hole faster. Tools amplify capability, they don't replace it.

I heard this exact same message over and over from hiring managers and engineering leaders at the Pragmatic Summit. Strong teams are getting stronger faster. Dysfunctional teams are getting dysfunctional faster. Some companies have cut customer-facing incidents in half since adopting AI tools. Others have doubled them. Same tools, completely different outcomes. The difference is the humans using them.

Now you might be asking yourself, "but what if AI gets WAY BETTER in 2–3 years? What if it can do the big-picture thinking too?"

Let's talk about what "better" actually means. Frontier model capabilities are absolutely still improving, but most of the improvements in performance that we're seeing aren't coming from bigger base models. They're coming from better tooling — things like improved context engineering and agent workflows. Understanding how to guide and improve agent systems will remain valuable skills for the foreseeable future.

And again, even if AI gets dramatically better at the "during code" phase, the verification, governance, communication, and accountability still require coding literacy.

Or what if you're thinking "I can vibe code apps without deep understanding already. Why bother learning?"

You can build demos and MVPs without really understanding how to code, sure. But production systems require a TON of things you don't know that you don't know if you've never learned this stuff from the ground up. If you want to ship something that handles real user data and real liability, you need to put the time in. Otherwise you're kind of stuck in Dunning-Kruger land.

And lastly, you might be aligned that learning all this makes sense on the job, but wondering if you can even get a job anyway. Junior hiring is really bad right now, so why even bother?

And yes, it's harder than 2021, no question about that. But it is still possible with the right projects, mindset, and strategy. I have tons of other videos on how to break in as a junior that I'll link below.

So if you're learning to code right now — or thinking about it — here's what I'd focus on. We can break this up into three steps:

### How to Learn in 2026

First, foundations. Pick one language and learn it really well. Python or JavaScript are good starting points. Understand fundamentals like data structures, APIs, authentication basics, and how databases work. Write unit tests and integration tests. And practice reading unfamiliar code and explaining what it does. This is the time to use AI only to explain concepts and test your understanding — don't outsource your learning to AI.

Once you've been studying for a while, ask yourself some questions: Can I read code and understand what it's doing? Can I debug a failing test? Can I reason about data flow and failure cases?

If yes, move on to step 2.

This is the "work with AI effectively" layer.

Learn to structure prompts with constraints and a clear definition of done. Use AI to generate tests, then audit them critically. Practice small, focused PRs instead of massive changes. Write evaluation checks for AI outputs. And treat code review as a primary skill.

Once you're confident that you can use AI to go faster without sacrificing correctness, you can move on to step 3.

This is the human layer, where you start practicing professional-level judgement.

Think about trade-offs of things like performance vs. cost, consistency vs. availability, or security and compliance. Write clear technical specs and design docs. Explain technical decisions to non-technical people — practice with your mom. Develop an incident response mindset: when things break, how do you triage and fix them?

Your goal should be to own a product end-to-end, from requirements to production.

I know that sounds like a lot. And it is! I'm not going to tell you this will be easy. And I'm not going to promise that if you learn to code, you'll definitely get a job.

The market is harder than it was a few years ago. AI is changing the way we work on a daily basis, and the skills that matter are changing too.

### So, Is Coding Dead?

You've probably heard some version of "coding is dead" recently. Maybe it was NVIDIA's CEO saying nobody will need to program anymore. Or Anthropic's CEO predicting AI would write 90% of code within 6 months (that was a year ago btw).

But like François Chollet, the creator of Keras, pointed out "software engineering has been within 6 months of being dead continually since early 2023."

And this pattern is way older than AI. FORTRAN was supposed to let scientists write programs without programmers. COBOL's English-like syntax was meant to let managers bypass developers entirely.

Every major abstraction — compilers, high-level languages, object-oriented programming — was pitched as making software engineers obsolete. But in reality, the demand for people who understand systems didn't disappear, it actually grew.

So you're not too late, and don't let the haters get you down.

— — —

If you're feeling like you need some support with your AI/ML career, here are some ways I can help:

- Subscribe to [my YouTube channel](https://www.youtube.com/@MarinaWyssAI) for weekly videos on technical topics, interviewing strategies, and more.
- Sign up for [my newsletter](https://www.gratitudedriven.com/subscribe) for a weekly post on a mix of technical topics and mindset/motivation for challenging fields.
- Want to level up your AI/ML career? Join the [AI/ML Career Launchpad](https://aiml-career-launchpad.circle.so/aiml-launchpad) community
- Interested in working with me 1:1? Learn more about my [strategic advisory sessions](https://www.marinawyss.com/coaching)