Calendar
===========

Tells when data was last updated in DBR


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
