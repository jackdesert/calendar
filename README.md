WhichFriday
===========

This is a calendar app, based on the design of pdxdailydancer.com
but this one instead of pulling from a mailing list,
it only displays recurring events.

Why?
----

Because I can never remember which event is on the third friday
of every month, much less figure out which friday is the third one.


Install Go
----------

http://golang.org/doc/install


Put This Source Code in the Appropriate Location
------------------------------------------------

The source code should live at

    $GOPATH/src/github.com/jackdesert/calendar

If it resides some other place, the import commands will complain that they can't find the packages.

Install Packages
----------------

We are using Godep in order to ensure that everyone running this software has the same version of each dependency.

    // ubuntu probably needs these packages
    sudo apt-get install mercurial bzr build-essential

    go get github.com/tools/godep
    godep restore


See https://github.com/tools/godep for more information on getting godep set up.


Test It
-------

Woops...no tests yet!


Run It Locally
--------------

    make && bin/calendar

Deploy
------

alias deploy_calendar="rsync -avr bin config README.md static index.html $dofreecinc:calendar && curl calendar.jackdesert.com:3501/restart"


Domain Names
------------

whichfriday.com
sexythursday.com
amazingcalendar.com
fuckingamazingcalendar.com
jalendar.com
favoritefriday.com


Features to Add
---------------

  * !!Sort events by start time
  * Teach it to respect central time
