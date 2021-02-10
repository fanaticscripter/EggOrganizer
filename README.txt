

                  ...
              .;loooool;.
            ,loooooooooool'
          .loooooooooooooool.
         ;ooooooooooooooooooo;              .cclllll;                         .;:,
        coooooooooooooooooooooc             ,X,......                        kd''cK'
       loooooooooooooooooooooooc            :K                              kd    kd
      coooooooooooooooooooooooooc           c0                             cK     dx
     :ooooooooooooooooooooooooooo;          l0                             Oo     dx                                        .
    'ooooooooooooooooooooooooooooo.         lKcccl:      .''...    .''..' .Xc     dx            .''..'                      l.
    ooooooooooooooooooooooooooooool         oO....     ,0l;;c0K  '0o;;cOX. X:     dk ,K'lxxc  .Oo;;cOX.  oxdxk,   'KokddO' .X' cddddxKk   .dxdxd. 'K'cdxo
   ;ooooooooooooooooooooooooooooooo,        ok        .K:    :X  0l    ,X. K:     dk :XO'  ;  Oo    .X'      lX.  .Xd   cK .X,     .xk.  c0'   dk ;X0;  ,.
   ooooooooooooooooooooooooooooooool        oO        ,X.    ,X..X'    'X. kd     kd :X.     .X;    .X,  .:dxkX.  .X.   ,X. X;    :0;   .Xkoodkd. :X'
  .ooooooooooooooooooooooooooooooooo.       :K        'X'    :X..X,    ;X. :K.   .X; :X       K:    'X' :0;  ,X.  .X.   .X. 0:   xk.    'X,..     ;X.
  .ooooooooooooooooooooooooooooooooo.       'X.        xk   .0X. oO.  .OX.  oO.  xk  ;X.      l0.  .kX' 0o .;kX.  .X.   .X' Oc .Oo       dO.   ,l ,X.
  .ooooooooooooooooooooooooooooooooo.       .kddddddl   :dddodX.  ;dddddX.   'dddc   ;k        ,ddddoX' .col;.:dl .d    .d. l, cxddddo.   'ldddl. ,k
   loooooooooooooooooooooooooooooooc                         ;X.       'X.                          .X'
   .ooooooooooooooooooooooooooooooo.                         ,X.       'X.                          .X,
    .ooooooooooooooooooooooooooool.                          cK        ;K.                          ,X.
     .cooooooooooooooooooooooooo:                      .,...:0;  .,...:0:                     .,...;Ol
       .coooooooooooooooooooooc.                       .;:c:,     ,:c:,.                       ,:c:,.
         .'coooooooooooooooc'.
             ..,;::c::;,..


Copyright (c) 2021 @mk2

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


Installation
============

Not necessary. You operate from this directory, so place this directory wherever you want.


Usage
=====

1. Edit the configuration file, config.toml, with a text editor. This is where you put information about players in your group.

2. Run the program. This is a command line program, so instructions differ for different operating systems.

   - macOS:

     - Launch Terminal.app from Spotlight.

     - Navigate to this directory with the `cd` command, e.g.

           cd ~/Downloads/EggOrganizer

     - Run

           .\EggOrganizer <arguments...>

   - Windows:

     - Launch PowerShell from Start Menu search.

     - Navigate to this directory with the `cd` command, e.g.

           cd D:\Downloads\EggOrganizer

     - Run

           .\EggOrganizer <arguments...>

   - Linux: Linux executable is not included. Ask the author.

Features and invocations
------------------------

- Check which players in the group have completed/attempted/never attempted a certain contract (all incarnations of it, including original and leggacy runs):

      EggOrganizer status <contract-id>

  E.g.

      EggOrganizer status space-tourism

  Use the -H, --hide-done flag to hide players who have completed the contract, e.g.

      EggOrganizer status -H space-tourism

  Use the -n, --no-id flag to remove the User ID column from output, useful for privacy-preservation when publishing the output:

      EggOrganizer status -n space-tourism

  Flags can be combined.
