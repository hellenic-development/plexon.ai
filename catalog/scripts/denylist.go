package catalogtypes

import "regexp"

// Vendor denylist.
//
// Some model vendors publish terms that reserve the right to train their
// general models on content submitted to them. Listing such a vendor's models
// in this catalog puts Plexon in conflict with the Limited Use requirement of
// the Google API Services User Data Policy, because a user can route Google
// user data (Gmail, Drive, Calendar) to whichever model they have selected.
//
// Google's Third Party Data Safety team rejected OAuth verification for the
// Cloud project `plexon-workspace-standard` on 2026-09-01 for exactly this
// reason, citing DeepSeek. The models it flagged were open-weight models served
// by third parties (Fireworks AI, OpenRouter) rather than a DeepSeek
// integration, but the listing itself was the finding, so the listing is gone.
//
// This check is what makes that removal durable rather than a one-time delete.
// It is enforced twice on purpose:
//
//   - sync_models_dev drops denied models on ingest, so the weekly models.dev
//     sync cannot reintroduce them, including ids that do not exist yet.
//   - build_snapshot fails outright if one is present, so a hand-edited YAML
//     can never reach snapshots/latest.json.
//
// Matching is on substring rather than exact id because vendors ship new ids
// constantly (deepseek-v4-flash-0731, deepseek/deepseek-v3.2-exp,
// ~deepseek/deepseek-v4-flash-latest), and an exact-id list goes stale the week
// after it is written. That is the failure mode SyncPolicy.Excluded has.
//
// DO NOT delete this to make a build pass. Removing a vendor here is a policy
// decision that has to be made deliberately and reviewed, and it needs the
// Google OAuth verification status for both Cloud projects re-checked first.
// See docs/mcp/GOOGLE_OAUTH_VERIFICATION.md in the private plexon repo, and
// docs/providers/DEEPSEEK_REMOVAL.md for the entries that were pulled.
var deniedVendorRe = regexp.MustCompile(`(?i)deep[ _-]*seek`)

// DeniedVendor reports whether a model belongs to a denied vendor, checking
// both its catalog id and its display name. Either matching is enough: a
// provider can namespace a model under its own prefix (accounts/fireworks/
// models/deepseek-v4-pro) while the display name carries the vendor, or the
// reverse.
func DeniedVendor(id, name string) bool {
	return deniedVendorRe.MatchString(id) || deniedVendorRe.MatchString(name)
}
