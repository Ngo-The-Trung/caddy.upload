// UploadHandler provides facilities to conduct uploads
// using the HTTP protocol as transport.
//
// If possible, i. e. if the operating- and filesystem implements it,
// files will not emerge before their first upload is completed.
// This is of importance to software that monitors a set of paths and
// reacts to new files. For example, with the intention to trigger uploads
// to other locations (mirrors).
//
// Requests are authenticated by sending a header "Authorization" formatted like this:
//
//  Authorization: Signature keyId="(key_id)",algorithm="hmac-sha256",
//      headers="timestamp token",signature="(see below)"
//
// The first element in 'headers' must either be "timestamp" (recommended),
// or "date" referring to HTTP header "Date".
// github.com/joyent/gosign is an implementation in Golang,
// github.com/joyent/node-http-signature for Node.js.
//
// This is how you generate said signature using a Linux shell:
//  key="geheim"
//  timestamp="$(date --utc +%s)"
//  token="streng"
//
//  printf "${timestamp}${token}" \
//  | openssl dgst -sha256 -hmac "${key}" -binary \
//  | openssl enc -base64
//
package upload // import "blitznote.com/src/caddy.upload"
