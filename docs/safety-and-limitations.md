# Safety and limitations

## Intended use

`ballistic_calc` estimates external-ballistic flight from user-supplied inputs
and a documented numerical model. Its results can support educational analysis,
comparison, planning, tables, charts, and DOPE-oriented views. They remain model
predictions, not observations or guarantees.

## Reusable product notice

The following concise language is suitable for the engine documentation,
standalone calculator, and KickBrass integration:

> External-ballistics results are estimates and depend on input quality, model
> limits, equipment and ammunition variation, weather, and numerical choices.
> They do not guarantee impacts, backstops, safe trajectories, or conditions.
> Verify results through appropriate real-world procedures and remain
> responsible for firearm safety and applicable law. This calculator does not
> provide reloading data, recommend loads, validate pressure, certify
> ammunition, or replace published manuals.

The full context below accompanies that short notice in product help and legal
review material.

## Technical limitations

External ballistics models a complex physical process using simplified inputs
and equations. Differences can come from:

- actual muzzle velocity and its variation;
- projectile shape, drag-coefficient variation, and ballistic-coefficient
  accuracy;
- firearm, sight, mount, bore, zero, and alignment differences;
- temperature, station pressure, humidity, altitude, and local weather
  measurement errors;
- wind variation over distance, terrain, and time;
- aerodynamic effects omitted from v0.1;
- numerical integration, interpolation, convergence, and rounding choices; and
- differences between the modeled atmosphere and the projectile's actual path.

The v0.1 point-mass model covers G1/G7 ballistic coefficient, level or inclined
shots, and no wind or one constant wind. It does not model range-varying wind,
cant, spin drift, Coriolis, aerodynamic jump, moving targets, every transonic
effect, weapon dynamics, terminal effects, or local hazards.

A close match to one observation does not establish accuracy for another
projectile, range, atmosphere, firearm, or wind condition. A displayed number
of decimal places does not imply matching physical certainty.

## Not a safety-critical firing system

The library is not designed or certified for safety-critical firing decisions.
It does not determine whether a target, firing direction, backstop, impact area,
or surrounding environment is safe. It does not account for every hazard,
ricochet, obstruction, equipment fault, bystander, legal restriction, or
changing condition.

Users independently verify zeros and predicted trajectories through appropriate
real-world procedures, use conservative safety practices, maintain a safe
backstop and firing environment, follow equipment and range instructions, and
comply with applicable law. The project does not provide legal advice.

## No reloading or ammunition-safety advice

This engine calculates external flight after launch. It does not:

- provide charge weights or load recipes;
- recommend powders, primers, projectiles, cases, seating depths, substitutions,
  or component combinations;
- calculate or validate chamber pressure;
- diagnose pressure from user observations;
- decide that a recipe, revision, batch, or ammo lot is safe;
- certify ammunition or firearm suitability; or
- replace current published manuals, manufacturer instructions, standards, or
  qualified professional guidance.

Projectile weight and muzzle velocity are physical inputs to a trajectory, not
a recommendation to assemble or fire ammunition with those values.

## Predicted and measured information

KickBrass and other consumers keep modeled results visibly distinct from:

- chronograph readings;
- observed impacts and group measurements;
- range-test notes;
- environmental measurements;
- pressure observations entered by a user; and
- recipe, batch, firearm, or component records.

A calculator result never overwrites measured data. A comparison can describe
agreement or difference, but it does not convert a prediction into a
measurement or a measurement into safety certification.

Bench prefill remains reviewable. Record completeness, catalog matching,
“active” or “tested” status, and a user's verdict do not silently validate a
calculation or imply ammunition safety.

## Error and presentation posture

Invalid, non-finite, unsupported, or non-convergent calculations produce an
explicit failure rather than plausible-looking partial output. Consumer
interfaces preserve units, input provenance, engine version, and calculation
time so a result can be interpreted in context.

Tables and charts label results as predicted or modeled. Warnings are concise
and placed at calculation entry, result interpretation, export, and saved
snapshot boundaries where context can otherwise be lost. The product avoids
both hidden disclaimers and repetitive alarm language.

## Release review

Numerical verification does not establish product safety, legal compliance, or
permission to distribute reference material. A public release of the engine or
KickBrass calculator requires separate provenance/legal review and product-
safety review, including the intended audience, claims, UI language, exports,
telemetry, retention, and integration behavior.

This document describes the project's intended boundary and is not a warranty,
certification, or legal opinion.
