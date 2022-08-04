// SPDX-License-Identifier: BSD-2-Clause
// Copyright (c) 2022-2022 Bitmark Inc.
// Use of this source code is governed by an BSD 2 Clause
// license that can be found in the LICENSE file.

// s3 client implements a simple client with upload struct as JSON record
// the JSON is '\n' terminated and is added in a path format base on a timestamp
// that is compatible with Athena database partitioned tables
package s3client
