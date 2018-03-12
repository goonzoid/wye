#!/usr/bin/env bats

go build -o wye

@test "mono files" {
  run ./wye fixtures/mono.wav
  [ "$status" -eq 0 ]
  [ "$output" = "this file is mono" ]
}

@test "stereo files with mono information" {
  run ./wye fixtures/fake_stereo.wav
  [ "$status" -eq 0 ]
  [ "$output" = "this file is stereo, but contains mono information" ]
}

@test "stereo files with stereo information" {
  run ./wye fixtures/true_stereo.wav
  [ "$status" -eq 0 ]
  [ "$output" = "this file is stereo, and contains stereo information" ]
}
