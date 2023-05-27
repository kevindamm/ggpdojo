Games
=====

Collection of game metadata and their definitions, where a game must consist of
two or more player agents; for single-player games see the `/puzzles` directory.

Contents
--------

**`README`** (required): Each directory must include at least a `README` or
`README.md` with the game description and its rules in a human-readable form.
The README may also include the state of the art for the game (whether it has
been solved, whether an AI made specifically for the game has achieved
tournament play, etc.).

Additional files in the game-specific directory include:

**`METADATA.json`**: Game metadata such as the number of players, its source and
author, various qualia about the game (its complexity, elegance, shibui, etc.)
and tags for collecting similar games.  Represented as a JSON dictionary whose
schmea can be found in `/schema/game_metadata.fbs`.

**`rulesheet.kif`**: Game definition in Game Description Language (GDL) in a
prefix-formtted KIF (as found in Stanford and Dresden repos, indeed some of the
initial game collection use games shared from there, see the game's metadata for
details).  This is the representation shared with clients (via gameplay API) and
is used as a compatibility layer for interplay with GGP competitions.

