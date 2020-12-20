Give the deserved importance to your commit messages
Why commit messages are important and why you should invest time on them
2020-12-20

Could you tell what is the most important part of your daily tasks as a Software Engineer ?
It is easy to think of delivering that bug fix on a good time manner, a feature or even
a preventive email to your team, one or another it all boils down to communication.

Communication exists in many forms in our daily tasks and its effectiveness reflects a team
performance, the reason is that Software Engineering for B2C, — but not only —, supports
the organization model provided by the company's product solution.

We rely on frameworks and concepts to communicate behaviour in the lowest level of computing
to the highest level of human interaction and if we do that badly we deteriorate productivity.
An easy spot to identify those degradations is on code delivery, Version Control Systems have
not only the responsibility to keep a history of changes but also to tell a history of evolution,
— or to better define —, to express the motivation behind an enterprise's product evolvement.

Mostly developers tend to not give the right importance to commit messages because
they care too much in coding and less about the business and the team  — don't be this developer.
This may work on startups or small code bases but on well distributed systems and teams, this is the
key to be unproductive.

Can we do better ?
Can we allow our team to become more productive by just commit messaging ?

Document motivation, dedicate time for telling the history you are sharing with the business of your company,
commit messages should reflect what motivated the changes you made and not what you changed
eg: add new json marshaling, extracted handler to new file, provided new list of customers.
By looking at changed files they already tell me that, so why be so redundant and unspecific about the changes ?
The mainly drawback of such messages is that they say nothing about the motivation, the product feature's state,
vision or current behaviour. things that are very important to understand how a
business works and why things works as they do.

Commit messages should be covered of references, motivation and answer questions like:
Why I am adding that ?
Which functionality is being affected and why ?
Does this change resulted from a previous change ? then reference it.

They should express a business intention behind the implementation, so how could we use the above example
to express our delivery ?

Instead of: "add new json marshaling"

Try:
"previously marshaling json data coming from redis were
causing float overflows, this commit decorates json marshaling
preventing data type overflows."

Instead of: "extracted handler to new file"

Try:  
"handler file were mixing more than enough of handlers,
this was causing maintenance confusion and breaking
single responsibility principle. this extract handlers to
its own file, create a handler type to better identify them implementing
http.Handler interface."

Instead of: "provide new list of customers"

Try:  
"customer list fix implemented by 30f960e was causing
the list to be unsorted, this fix the sort using the customer's
last name."

Truth be told, it's not an easy task for a developer who just wants to delivery a feature document its changes in this way,
the pressure and the will to delivery speaks higher, and the quality of the message degrades as he/she attempts to pass
the build. Give it a try, document the motivation, spread the word to your team, and you will see how positively this can
affect the overall productivity.

The taste of digging deeper code bases trying to understand why the previous engineer did what is seen while saying
"add new discount rule" will not be bitter anymore and will allow the reason be directed to the domain owners in a more
confident and quicker way than ever.






























