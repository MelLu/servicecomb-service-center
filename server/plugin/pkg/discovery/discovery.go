/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package discovery

// Adaptor is used to do service discovery.
// To improve the performance, Adaptor may use cache firstly in
// service discovery.
type Adaptor interface {
	Runnable
	// Indexer is used to search data from the registry.
	Indexer
	// Cacher is used to manage the registry's cache.
	Cacher
}

// AdaptorRepository creates Adaptors
type AdaptorRepository interface {
	// New news an instance of specify Type adaptor
	New(t Type, cfg *Config) Adaptor
}
