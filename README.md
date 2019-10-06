endsentence
================

End your comments with a period. This will check your go files for comments on
methods/functions that are commented but does not have a period in the end of
the sentence.

Exceptions
------------

There are a few minor exceptions to when endsentence does not require a period
at the end of a sentence:

* If your comment ends with a URL or an email
* If your comment ends with a list (starts with a `*` or `-`)

Neomake
------------

To add this to Vim and Neomake, use the following configuration:

    let g:neomake_go_endsentence_maker = {
       \ 'args': [
       \   '.',
       \ ],
       \ 'append_file': 0,
       \ 'cwd': '%:p:h',
       \ }

Reporting bugs
--------------

If you find any bugs or want to provide feedback, you can file bugs in the project's [GitHub Issues page](https://github.com/antonlindstrom/endsentence).

Author
------

This project is maintained by [Anton Lindström](https://www.antonlindstrom.com) ([GitHub](https://github.com/antonlindstrom) | [Twitter](https://twitter.com/mycap))

License
-------

APACHE LICENSE 2.0
Copyright 2019 Anton Lindström

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
