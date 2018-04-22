package cornerdetect

import (
	"image"
)

// fast9_corner_score - transpiled function from  fast_9.c:7
// This is mechanically generated code
// Compute the score using binary search
func fast9_corner_score(im []uint8, offset int, pixel []int, bstart int) (c2goDefaultReturn int) {
	var bmin = bstart
	var bmax = 255
	var b = (bmax + bmin) / 2
	for {
		var cb = int(im[offset]) + (b)
		var c_b = int(im[offset]) - (b)
		if int(im[pixel[0]+offset]) > cb {
			if int(im[pixel[1]+offset]) > cb {
				if int(im[pixel[2]+offset]) > cb {
					if int(im[pixel[3]+offset]) > cb {
						if int(im[pixel[4]+offset]) > cb {
							if int(im[pixel[5]+offset]) > cb {
								if int(im[pixel[6]+offset]) > cb {
									if int(im[pixel[7]+offset]) > cb {
										if int(im[pixel[8]+offset]) > cb {
											goto is_a_corner
										} else {
											if int(im[pixel[15]+offset]) > cb {
												goto is_a_corner
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										if int(im[pixel[7]+offset]) < c_b {
											if int(im[pixel[14]+offset]) > cb {
												if int(im[pixel[15]+offset]) > cb {
													goto is_a_corner
												} else {
													goto is_not_a_corner
												}
											} else {
												if int(im[pixel[14]+offset]) < c_b {
													if int(im[pixel[8]+offset]) < c_b {
														if int(im[pixel[9]+offset]) < c_b {
															if int(im[pixel[10]+offset]) < c_b {
																if int(im[pixel[11]+offset]) < c_b {
																	if int(im[pixel[12]+offset]) < c_b {
																		if int(im[pixel[13]+offset]) < c_b {
																			if int(im[pixel[15]+offset]) < c_b {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[14]+offset]) > cb {
												if int(im[pixel[15]+offset]) > cb {
													goto is_a_corner
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									}
								} else {
									if int(im[pixel[6]+offset]) < c_b {
										if int(im[pixel[15]+offset]) > cb {
											if int(im[pixel[13]+offset]) > cb {
												if int(im[pixel[14]+offset]) > cb {
													goto is_a_corner
												} else {
													goto is_not_a_corner
												}
											} else {
												if int(im[pixel[13]+offset]) < c_b {
													if int(im[pixel[7]+offset]) < c_b {
														if int(im[pixel[8]+offset]) < c_b {
															if int(im[pixel[9]+offset]) < c_b {
																if int(im[pixel[10]+offset]) < c_b {
																	if int(im[pixel[11]+offset]) < c_b {
																		if int(im[pixel[12]+offset]) < c_b {
																			if int(im[pixel[14]+offset]) < c_b {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[7]+offset]) < c_b {
												if int(im[pixel[8]+offset]) < c_b {
													if int(im[pixel[9]+offset]) < c_b {
														if int(im[pixel[10]+offset]) < c_b {
															if int(im[pixel[11]+offset]) < c_b {
																if int(im[pixel[12]+offset]) < c_b {
																	if int(im[pixel[13]+offset]) < c_b {
																		if int(im[pixel[14]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										if int(im[pixel[13]+offset]) > cb {
											if int(im[pixel[14]+offset]) > cb {
												if int(im[pixel[15]+offset]) > cb {
													goto is_a_corner
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											if int(im[pixel[13]+offset]) < c_b {
												if int(im[pixel[7]+offset]) < c_b {
													if int(im[pixel[8]+offset]) < c_b {
														if int(im[pixel[9]+offset]) < c_b {
															if int(im[pixel[10]+offset]) < c_b {
																if int(im[pixel[11]+offset]) < c_b {
																	if int(im[pixel[12]+offset]) < c_b {
																		if int(im[pixel[14]+offset]) < c_b {
																			if int(im[pixel[15]+offset]) < c_b {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									}
								}
							} else {
								if int(im[pixel[5]+offset]) < c_b {
									if int(im[pixel[14]+offset]) > cb {
										if int(im[pixel[12]+offset]) > cb {
											if int(im[pixel[13]+offset]) > cb {
												if int(im[pixel[15]+offset]) > cb {
													goto is_a_corner
												} else {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	if int(im[pixel[10]+offset]) > cb {
																		if int(im[pixel[11]+offset]) > cb {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											if int(im[pixel[12]+offset]) < c_b {
												if int(im[pixel[6]+offset]) < c_b {
													if int(im[pixel[7]+offset]) < c_b {
														if int(im[pixel[8]+offset]) < c_b {
															if int(im[pixel[9]+offset]) < c_b {
																if int(im[pixel[10]+offset]) < c_b {
																	if int(im[pixel[11]+offset]) < c_b {
																		if int(im[pixel[13]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										if int(im[pixel[14]+offset]) < c_b {
											if int(im[pixel[7]+offset]) < c_b {
												if int(im[pixel[8]+offset]) < c_b {
													if int(im[pixel[9]+offset]) < c_b {
														if int(im[pixel[10]+offset]) < c_b {
															if int(im[pixel[11]+offset]) < c_b {
																if int(im[pixel[12]+offset]) < c_b {
																	if int(im[pixel[13]+offset]) < c_b {
																		if int(im[pixel[6]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			if int(im[pixel[15]+offset]) < c_b {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											if int(im[pixel[6]+offset]) < c_b {
												if int(im[pixel[7]+offset]) < c_b {
													if int(im[pixel[8]+offset]) < c_b {
														if int(im[pixel[9]+offset]) < c_b {
															if int(im[pixel[10]+offset]) < c_b {
																if int(im[pixel[11]+offset]) < c_b {
																	if int(im[pixel[12]+offset]) < c_b {
																		if int(im[pixel[13]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									}
								} else {
									if int(im[pixel[12]+offset]) > cb {
										if int(im[pixel[13]+offset]) > cb {
											if int(im[pixel[14]+offset]) > cb {
												if int(im[pixel[15]+offset]) > cb {
													goto is_a_corner
												} else {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	if int(im[pixel[10]+offset]) > cb {
																		if int(im[pixel[11]+offset]) > cb {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										if int(im[pixel[12]+offset]) < c_b {
											if int(im[pixel[7]+offset]) < c_b {
												if int(im[pixel[8]+offset]) < c_b {
													if int(im[pixel[9]+offset]) < c_b {
														if int(im[pixel[10]+offset]) < c_b {
															if int(im[pixel[11]+offset]) < c_b {
																if int(im[pixel[13]+offset]) < c_b {
																	if int(im[pixel[14]+offset]) < c_b {
																		if int(im[pixel[6]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			if int(im[pixel[15]+offset]) < c_b {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									}
								}
							}
						} else {
							if int(im[pixel[4]+offset]) < c_b {
								if int(im[pixel[13]+offset]) > cb {
									if int(im[pixel[11]+offset]) > cb {
										if int(im[pixel[12]+offset]) > cb {
											if int(im[pixel[14]+offset]) > cb {
												if int(im[pixel[15]+offset]) > cb {
													goto is_a_corner
												} else {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	if int(im[pixel[10]+offset]) > cb {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	if int(im[pixel[10]+offset]) > cb {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										if int(im[pixel[11]+offset]) < c_b {
											if int(im[pixel[5]+offset]) < c_b {
												if int(im[pixel[6]+offset]) < c_b {
													if int(im[pixel[7]+offset]) < c_b {
														if int(im[pixel[8]+offset]) < c_b {
															if int(im[pixel[9]+offset]) < c_b {
																if int(im[pixel[10]+offset]) < c_b {
																	if int(im[pixel[12]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									}
								} else {
									if int(im[pixel[13]+offset]) < c_b {
										if int(im[pixel[7]+offset]) < c_b {
											if int(im[pixel[8]+offset]) < c_b {
												if int(im[pixel[9]+offset]) < c_b {
													if int(im[pixel[10]+offset]) < c_b {
														if int(im[pixel[11]+offset]) < c_b {
															if int(im[pixel[12]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[5]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		if int(im[pixel[14]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	}
																} else {
																	if int(im[pixel[14]+offset]) < c_b {
																		if int(im[pixel[15]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										if int(im[pixel[5]+offset]) < c_b {
											if int(im[pixel[6]+offset]) < c_b {
												if int(im[pixel[7]+offset]) < c_b {
													if int(im[pixel[8]+offset]) < c_b {
														if int(im[pixel[9]+offset]) < c_b {
															if int(im[pixel[10]+offset]) < c_b {
																if int(im[pixel[11]+offset]) < c_b {
																	if int(im[pixel[12]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									}
								}
							} else {
								if int(im[pixel[11]+offset]) > cb {
									if int(im[pixel[12]+offset]) > cb {
										if int(im[pixel[13]+offset]) > cb {
											if int(im[pixel[14]+offset]) > cb {
												if int(im[pixel[15]+offset]) > cb {
													goto is_a_corner
												} else {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	if int(im[pixel[10]+offset]) > cb {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	if int(im[pixel[10]+offset]) > cb {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									if int(im[pixel[11]+offset]) < c_b {
										if int(im[pixel[7]+offset]) < c_b {
											if int(im[pixel[8]+offset]) < c_b {
												if int(im[pixel[9]+offset]) < c_b {
													if int(im[pixel[10]+offset]) < c_b {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[13]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[5]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		if int(im[pixel[14]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	}
																} else {
																	if int(im[pixel[14]+offset]) < c_b {
																		if int(im[pixel[15]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								}
							}
						}
					} else {
						if int(im[pixel[3]+offset]) < c_b {
							if int(im[pixel[10]+offset]) > cb {
								if int(im[pixel[11]+offset]) > cb {
									if int(im[pixel[12]+offset]) > cb {
										if int(im[pixel[13]+offset]) > cb {
											if int(im[pixel[14]+offset]) > cb {
												if int(im[pixel[15]+offset]) > cb {
													goto is_a_corner
												} else {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									goto is_not_a_corner
								}
							} else {
								if int(im[pixel[10]+offset]) < c_b {
									if int(im[pixel[7]+offset]) < c_b {
										if int(im[pixel[8]+offset]) < c_b {
											if int(im[pixel[9]+offset]) < c_b {
												if int(im[pixel[11]+offset]) < c_b {
													if int(im[pixel[6]+offset]) < c_b {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[4]+offset]) < c_b {
																goto is_a_corner
															} else {
																if int(im[pixel[12]+offset]) < c_b {
																	if int(im[pixel[13]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															}
														} else {
															if int(im[pixel[12]+offset]) < c_b {
																if int(im[pixel[13]+offset]) < c_b {
																	if int(im[pixel[14]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[13]+offset]) < c_b {
																if int(im[pixel[14]+offset]) < c_b {
																	if int(im[pixel[15]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									goto is_not_a_corner
								}
							}
						} else {
							if int(im[pixel[10]+offset]) > cb {
								if int(im[pixel[11]+offset]) > cb {
									if int(im[pixel[12]+offset]) > cb {
										if int(im[pixel[13]+offset]) > cb {
											if int(im[pixel[14]+offset]) > cb {
												if int(im[pixel[15]+offset]) > cb {
													goto is_a_corner
												} else {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									goto is_not_a_corner
								}
							} else {
								if int(im[pixel[10]+offset]) < c_b {
									if int(im[pixel[7]+offset]) < c_b {
										if int(im[pixel[8]+offset]) < c_b {
											if int(im[pixel[9]+offset]) < c_b {
												if int(im[pixel[11]+offset]) < c_b {
													if int(im[pixel[12]+offset]) < c_b {
														if int(im[pixel[6]+offset]) < c_b {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[4]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	if int(im[pixel[13]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																}
															} else {
																if int(im[pixel[13]+offset]) < c_b {
																	if int(im[pixel[14]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															}
														} else {
															if int(im[pixel[13]+offset]) < c_b {
																if int(im[pixel[14]+offset]) < c_b {
																	if int(im[pixel[15]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									goto is_not_a_corner
								}
							}
						}
					}
				} else {
					if int(im[pixel[2]+offset]) < c_b {
						if int(im[pixel[9]+offset]) > cb {
							if int(im[pixel[10]+offset]) > cb {
								if int(im[pixel[11]+offset]) > cb {
									if int(im[pixel[12]+offset]) > cb {
										if int(im[pixel[13]+offset]) > cb {
											if int(im[pixel[14]+offset]) > cb {
												if int(im[pixel[15]+offset]) > cb {
													goto is_a_corner
												} else {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										if int(im[pixel[3]+offset]) > cb {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									}
								} else {
									goto is_not_a_corner
								}
							} else {
								goto is_not_a_corner
							}
						} else {
							if int(im[pixel[9]+offset]) < c_b {
								if int(im[pixel[7]+offset]) < c_b {
									if int(im[pixel[8]+offset]) < c_b {
										if int(im[pixel[10]+offset]) < c_b {
											if int(im[pixel[6]+offset]) < c_b {
												if int(im[pixel[5]+offset]) < c_b {
													if int(im[pixel[4]+offset]) < c_b {
														if int(im[pixel[3]+offset]) < c_b {
															goto is_a_corner
														} else {
															if int(im[pixel[11]+offset]) < c_b {
																if int(im[pixel[12]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														if int(im[pixel[11]+offset]) < c_b {
															if int(im[pixel[12]+offset]) < c_b {
																if int(im[pixel[13]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													if int(im[pixel[11]+offset]) < c_b {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[13]+offset]) < c_b {
																if int(im[pixel[14]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[11]+offset]) < c_b {
													if int(im[pixel[12]+offset]) < c_b {
														if int(im[pixel[13]+offset]) < c_b {
															if int(im[pixel[14]+offset]) < c_b {
																if int(im[pixel[15]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									goto is_not_a_corner
								}
							} else {
								goto is_not_a_corner
							}
						}
					} else {
						if int(im[pixel[9]+offset]) > cb {
							if int(im[pixel[10]+offset]) > cb {
								if int(im[pixel[11]+offset]) > cb {
									if int(im[pixel[12]+offset]) > cb {
										if int(im[pixel[13]+offset]) > cb {
											if int(im[pixel[14]+offset]) > cb {
												if int(im[pixel[15]+offset]) > cb {
													goto is_a_corner
												} else {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										if int(im[pixel[3]+offset]) > cb {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									}
								} else {
									goto is_not_a_corner
								}
							} else {
								goto is_not_a_corner
							}
						} else {
							if int(im[pixel[9]+offset]) < c_b {
								if int(im[pixel[7]+offset]) < c_b {
									if int(im[pixel[8]+offset]) < c_b {
										if int(im[pixel[10]+offset]) < c_b {
											if int(im[pixel[11]+offset]) < c_b {
												if int(im[pixel[6]+offset]) < c_b {
													if int(im[pixel[5]+offset]) < c_b {
														if int(im[pixel[4]+offset]) < c_b {
															if int(im[pixel[3]+offset]) < c_b {
																goto is_a_corner
															} else {
																if int(im[pixel[12]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															}
														} else {
															if int(im[pixel[12]+offset]) < c_b {
																if int(im[pixel[13]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[13]+offset]) < c_b {
																if int(im[pixel[14]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													if int(im[pixel[12]+offset]) < c_b {
														if int(im[pixel[13]+offset]) < c_b {
															if int(im[pixel[14]+offset]) < c_b {
																if int(im[pixel[15]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									goto is_not_a_corner
								}
							} else {
								goto is_not_a_corner
							}
						}
					}
				}
			} else {
				if int(im[pixel[1]+offset]) < c_b {
					if int(im[pixel[8]+offset]) > cb {
						if int(im[pixel[9]+offset]) > cb {
							if int(im[pixel[10]+offset]) > cb {
								if int(im[pixel[11]+offset]) > cb {
									if int(im[pixel[12]+offset]) > cb {
										if int(im[pixel[13]+offset]) > cb {
											if int(im[pixel[14]+offset]) > cb {
												if int(im[pixel[15]+offset]) > cb {
													goto is_a_corner
												} else {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										if int(im[pixel[3]+offset]) > cb {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									}
								} else {
									if int(im[pixel[2]+offset]) > cb {
										if int(im[pixel[3]+offset]) > cb {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								}
							} else {
								goto is_not_a_corner
							}
						} else {
							goto is_not_a_corner
						}
					} else {
						if int(im[pixel[8]+offset]) < c_b {
							if int(im[pixel[7]+offset]) < c_b {
								if int(im[pixel[9]+offset]) < c_b {
									if int(im[pixel[6]+offset]) < c_b {
										if int(im[pixel[5]+offset]) < c_b {
											if int(im[pixel[4]+offset]) < c_b {
												if int(im[pixel[3]+offset]) < c_b {
													if int(im[pixel[2]+offset]) < c_b {
														goto is_a_corner
													} else {
														if int(im[pixel[10]+offset]) < c_b {
															if int(im[pixel[11]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													if int(im[pixel[10]+offset]) < c_b {
														if int(im[pixel[11]+offset]) < c_b {
															if int(im[pixel[12]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[10]+offset]) < c_b {
													if int(im[pixel[11]+offset]) < c_b {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[13]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[10]+offset]) < c_b {
												if int(im[pixel[11]+offset]) < c_b {
													if int(im[pixel[12]+offset]) < c_b {
														if int(im[pixel[13]+offset]) < c_b {
															if int(im[pixel[14]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										if int(im[pixel[10]+offset]) < c_b {
											if int(im[pixel[11]+offset]) < c_b {
												if int(im[pixel[12]+offset]) < c_b {
													if int(im[pixel[13]+offset]) < c_b {
														if int(im[pixel[14]+offset]) < c_b {
															if int(im[pixel[15]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									}
								} else {
									goto is_not_a_corner
								}
							} else {
								goto is_not_a_corner
							}
						} else {
							goto is_not_a_corner
						}
					}
				} else {
					if int(im[pixel[8]+offset]) > cb {
						if int(im[pixel[9]+offset]) > cb {
							if int(im[pixel[10]+offset]) > cb {
								if int(im[pixel[11]+offset]) > cb {
									if int(im[pixel[12]+offset]) > cb {
										if int(im[pixel[13]+offset]) > cb {
											if int(im[pixel[14]+offset]) > cb {
												if int(im[pixel[15]+offset]) > cb {
													goto is_a_corner
												} else {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										if int(im[pixel[3]+offset]) > cb {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									}
								} else {
									if int(im[pixel[2]+offset]) > cb {
										if int(im[pixel[3]+offset]) > cb {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								}
							} else {
								goto is_not_a_corner
							}
						} else {
							goto is_not_a_corner
						}
					} else {
						if int(im[pixel[8]+offset]) < c_b {
							if int(im[pixel[7]+offset]) < c_b {
								if int(im[pixel[9]+offset]) < c_b {
									if int(im[pixel[10]+offset]) < c_b {
										if int(im[pixel[6]+offset]) < c_b {
											if int(im[pixel[5]+offset]) < c_b {
												if int(im[pixel[4]+offset]) < c_b {
													if int(im[pixel[3]+offset]) < c_b {
														if int(im[pixel[2]+offset]) < c_b {
															goto is_a_corner
														} else {
															if int(im[pixel[11]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														if int(im[pixel[11]+offset]) < c_b {
															if int(im[pixel[12]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													if int(im[pixel[11]+offset]) < c_b {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[13]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[11]+offset]) < c_b {
													if int(im[pixel[12]+offset]) < c_b {
														if int(im[pixel[13]+offset]) < c_b {
															if int(im[pixel[14]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[11]+offset]) < c_b {
												if int(im[pixel[12]+offset]) < c_b {
													if int(im[pixel[13]+offset]) < c_b {
														if int(im[pixel[14]+offset]) < c_b {
															if int(im[pixel[15]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									goto is_not_a_corner
								}
							} else {
								goto is_not_a_corner
							}
						} else {
							goto is_not_a_corner
						}
					}
				}
			}
		} else {
			if int(im[pixel[0]+offset]) < c_b {
				if int(im[pixel[1]+offset]) > cb {
					if int(im[pixel[8]+offset]) > cb {
						if int(im[pixel[7]+offset]) > cb {
							if int(im[pixel[9]+offset]) > cb {
								if int(im[pixel[6]+offset]) > cb {
									if int(im[pixel[5]+offset]) > cb {
										if int(im[pixel[4]+offset]) > cb {
											if int(im[pixel[3]+offset]) > cb {
												if int(im[pixel[2]+offset]) > cb {
													goto is_a_corner
												} else {
													if int(im[pixel[10]+offset]) > cb {
														if int(im[pixel[11]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[10]+offset]) > cb {
													if int(im[pixel[11]+offset]) > cb {
														if int(im[pixel[12]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[10]+offset]) > cb {
												if int(im[pixel[11]+offset]) > cb {
													if int(im[pixel[12]+offset]) > cb {
														if int(im[pixel[13]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										if int(im[pixel[10]+offset]) > cb {
											if int(im[pixel[11]+offset]) > cb {
												if int(im[pixel[12]+offset]) > cb {
													if int(im[pixel[13]+offset]) > cb {
														if int(im[pixel[14]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									}
								} else {
									if int(im[pixel[10]+offset]) > cb {
										if int(im[pixel[11]+offset]) > cb {
											if int(im[pixel[12]+offset]) > cb {
												if int(im[pixel[13]+offset]) > cb {
													if int(im[pixel[14]+offset]) > cb {
														if int(im[pixel[15]+offset]) > cb {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								}
							} else {
								goto is_not_a_corner
							}
						} else {
							goto is_not_a_corner
						}
					} else {
						if int(im[pixel[8]+offset]) < c_b {
							if int(im[pixel[9]+offset]) < c_b {
								if int(im[pixel[10]+offset]) < c_b {
									if int(im[pixel[11]+offset]) < c_b {
										if int(im[pixel[12]+offset]) < c_b {
											if int(im[pixel[13]+offset]) < c_b {
												if int(im[pixel[14]+offset]) < c_b {
													if int(im[pixel[15]+offset]) < c_b {
														goto is_a_corner
													} else {
														if int(im[pixel[6]+offset]) < c_b {
															if int(im[pixel[7]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													if int(im[pixel[5]+offset]) < c_b {
														if int(im[pixel[6]+offset]) < c_b {
															if int(im[pixel[7]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[4]+offset]) < c_b {
													if int(im[pixel[5]+offset]) < c_b {
														if int(im[pixel[6]+offset]) < c_b {
															if int(im[pixel[7]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[3]+offset]) < c_b {
												if int(im[pixel[4]+offset]) < c_b {
													if int(im[pixel[5]+offset]) < c_b {
														if int(im[pixel[6]+offset]) < c_b {
															if int(im[pixel[7]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										if int(im[pixel[2]+offset]) < c_b {
											if int(im[pixel[3]+offset]) < c_b {
												if int(im[pixel[4]+offset]) < c_b {
													if int(im[pixel[5]+offset]) < c_b {
														if int(im[pixel[6]+offset]) < c_b {
															if int(im[pixel[7]+offset]) < c_b {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									}
								} else {
									goto is_not_a_corner
								}
							} else {
								goto is_not_a_corner
							}
						} else {
							goto is_not_a_corner
						}
					}
				} else {
					if int(im[pixel[1]+offset]) < c_b {
						if int(im[pixel[2]+offset]) > cb {
							if int(im[pixel[9]+offset]) > cb {
								if int(im[pixel[7]+offset]) > cb {
									if int(im[pixel[8]+offset]) > cb {
										if int(im[pixel[10]+offset]) > cb {
											if int(im[pixel[6]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[4]+offset]) > cb {
														if int(im[pixel[3]+offset]) > cb {
															goto is_a_corner
														} else {
															if int(im[pixel[11]+offset]) > cb {
																if int(im[pixel[12]+offset]) > cb {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														if int(im[pixel[11]+offset]) > cb {
															if int(im[pixel[12]+offset]) > cb {
																if int(im[pixel[13]+offset]) > cb {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													if int(im[pixel[11]+offset]) > cb {
														if int(im[pixel[12]+offset]) > cb {
															if int(im[pixel[13]+offset]) > cb {
																if int(im[pixel[14]+offset]) > cb {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[11]+offset]) > cb {
													if int(im[pixel[12]+offset]) > cb {
														if int(im[pixel[13]+offset]) > cb {
															if int(im[pixel[14]+offset]) > cb {
																if int(im[pixel[15]+offset]) > cb {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									goto is_not_a_corner
								}
							} else {
								if int(im[pixel[9]+offset]) < c_b {
									if int(im[pixel[10]+offset]) < c_b {
										if int(im[pixel[11]+offset]) < c_b {
											if int(im[pixel[12]+offset]) < c_b {
												if int(im[pixel[13]+offset]) < c_b {
													if int(im[pixel[14]+offset]) < c_b {
														if int(im[pixel[15]+offset]) < c_b {
															goto is_a_corner
														} else {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																	if int(im[pixel[8]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																	if int(im[pixel[8]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													if int(im[pixel[4]+offset]) < c_b {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																	if int(im[pixel[8]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[3]+offset]) < c_b {
													if int(im[pixel[4]+offset]) < c_b {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																	if int(im[pixel[8]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									goto is_not_a_corner
								}
							}
						} else {
							if int(im[pixel[2]+offset]) < c_b {
								if int(im[pixel[3]+offset]) > cb {
									if int(im[pixel[10]+offset]) > cb {
										if int(im[pixel[7]+offset]) > cb {
											if int(im[pixel[8]+offset]) > cb {
												if int(im[pixel[9]+offset]) > cb {
													if int(im[pixel[11]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[5]+offset]) > cb {
																if int(im[pixel[4]+offset]) > cb {
																	goto is_a_corner
																} else {
																	if int(im[pixel[12]+offset]) > cb {
																		if int(im[pixel[13]+offset]) > cb {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																}
															} else {
																if int(im[pixel[12]+offset]) > cb {
																	if int(im[pixel[13]+offset]) > cb {
																		if int(im[pixel[14]+offset]) > cb {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															}
														} else {
															if int(im[pixel[12]+offset]) > cb {
																if int(im[pixel[13]+offset]) > cb {
																	if int(im[pixel[14]+offset]) > cb {
																		if int(im[pixel[15]+offset]) > cb {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										if int(im[pixel[10]+offset]) < c_b {
											if int(im[pixel[11]+offset]) < c_b {
												if int(im[pixel[12]+offset]) < c_b {
													if int(im[pixel[13]+offset]) < c_b {
														if int(im[pixel[14]+offset]) < c_b {
															if int(im[pixel[15]+offset]) < c_b {
																goto is_a_corner
															} else {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																		if int(im[pixel[8]+offset]) < c_b {
																			if int(im[pixel[9]+offset]) < c_b {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															}
														} else {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																		if int(im[pixel[8]+offset]) < c_b {
																			if int(im[pixel[9]+offset]) < c_b {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														if int(im[pixel[4]+offset]) < c_b {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																		if int(im[pixel[8]+offset]) < c_b {
																			if int(im[pixel[9]+offset]) < c_b {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									}
								} else {
									if int(im[pixel[3]+offset]) < c_b {
										if int(im[pixel[4]+offset]) > cb {
											if int(im[pixel[13]+offset]) > cb {
												if int(im[pixel[7]+offset]) > cb {
													if int(im[pixel[8]+offset]) > cb {
														if int(im[pixel[9]+offset]) > cb {
															if int(im[pixel[10]+offset]) > cb {
																if int(im[pixel[11]+offset]) > cb {
																	if int(im[pixel[12]+offset]) > cb {
																		if int(im[pixel[6]+offset]) > cb {
																			if int(im[pixel[5]+offset]) > cb {
																				goto is_a_corner
																			} else {
																				if int(im[pixel[14]+offset]) > cb {
																					goto is_a_corner
																				} else {
																					goto is_not_a_corner
																				}
																			}
																		} else {
																			if int(im[pixel[14]+offset]) > cb {
																				if int(im[pixel[15]+offset]) > cb {
																					goto is_a_corner
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												if int(im[pixel[13]+offset]) < c_b {
													if int(im[pixel[11]+offset]) > cb {
														if int(im[pixel[5]+offset]) > cb {
															if int(im[pixel[6]+offset]) > cb {
																if int(im[pixel[7]+offset]) > cb {
																	if int(im[pixel[8]+offset]) > cb {
																		if int(im[pixel[9]+offset]) > cb {
																			if int(im[pixel[10]+offset]) > cb {
																				if int(im[pixel[12]+offset]) > cb {
																					goto is_a_corner
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														if int(im[pixel[11]+offset]) < c_b {
															if int(im[pixel[12]+offset]) < c_b {
																if int(im[pixel[14]+offset]) < c_b {
																	if int(im[pixel[15]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		if int(im[pixel[6]+offset]) < c_b {
																			if int(im[pixel[7]+offset]) < c_b {
																				if int(im[pixel[8]+offset]) < c_b {
																					if int(im[pixel[9]+offset]) < c_b {
																						if int(im[pixel[10]+offset]) < c_b {
																							goto is_a_corner
																						} else {
																							goto is_not_a_corner
																						}
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	}
																} else {
																	if int(im[pixel[5]+offset]) < c_b {
																		if int(im[pixel[6]+offset]) < c_b {
																			if int(im[pixel[7]+offset]) < c_b {
																				if int(im[pixel[8]+offset]) < c_b {
																					if int(im[pixel[9]+offset]) < c_b {
																						if int(im[pixel[10]+offset]) < c_b {
																							goto is_a_corner
																						} else {
																							goto is_not_a_corner
																						}
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																		if int(im[pixel[10]+offset]) > cb {
																			if int(im[pixel[11]+offset]) > cb {
																				if int(im[pixel[12]+offset]) > cb {
																					goto is_a_corner
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											}
										} else {
											if int(im[pixel[4]+offset]) < c_b {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[14]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	if int(im[pixel[10]+offset]) > cb {
																		if int(im[pixel[11]+offset]) > cb {
																			if int(im[pixel[12]+offset]) > cb {
																				if int(im[pixel[13]+offset]) > cb {
																					if int(im[pixel[6]+offset]) > cb {
																						goto is_a_corner
																					} else {
																						if int(im[pixel[15]+offset]) > cb {
																							goto is_a_corner
																						} else {
																							goto is_not_a_corner
																						}
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														if int(im[pixel[14]+offset]) < c_b {
															if int(im[pixel[12]+offset]) > cb {
																if int(im[pixel[6]+offset]) > cb {
																	if int(im[pixel[7]+offset]) > cb {
																		if int(im[pixel[8]+offset]) > cb {
																			if int(im[pixel[9]+offset]) > cb {
																				if int(im[pixel[10]+offset]) > cb {
																					if int(im[pixel[11]+offset]) > cb {
																						if int(im[pixel[13]+offset]) > cb {
																							goto is_a_corner
																						} else {
																							goto is_not_a_corner
																						}
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																if int(im[pixel[12]+offset]) < c_b {
																	if int(im[pixel[13]+offset]) < c_b {
																		if int(im[pixel[15]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			if int(im[pixel[6]+offset]) < c_b {
																				if int(im[pixel[7]+offset]) < c_b {
																					if int(im[pixel[8]+offset]) < c_b {
																						if int(im[pixel[9]+offset]) < c_b {
																							if int(im[pixel[10]+offset]) < c_b {
																								if int(im[pixel[11]+offset]) < c_b {
																									goto is_a_corner
																								} else {
																									goto is_not_a_corner
																								}
																							} else {
																								goto is_not_a_corner
																							}
																						} else {
																							goto is_not_a_corner
																						}
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															}
														} else {
															if int(im[pixel[6]+offset]) > cb {
																if int(im[pixel[7]+offset]) > cb {
																	if int(im[pixel[8]+offset]) > cb {
																		if int(im[pixel[9]+offset]) > cb {
																			if int(im[pixel[10]+offset]) > cb {
																				if int(im[pixel[11]+offset]) > cb {
																					if int(im[pixel[12]+offset]) > cb {
																						if int(im[pixel[13]+offset]) > cb {
																							goto is_a_corner
																						} else {
																							goto is_not_a_corner
																						}
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													}
												} else {
													if int(im[pixel[5]+offset]) < c_b {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[15]+offset]) < c_b {
																if int(im[pixel[13]+offset]) > cb {
																	if int(im[pixel[7]+offset]) > cb {
																		if int(im[pixel[8]+offset]) > cb {
																			if int(im[pixel[9]+offset]) > cb {
																				if int(im[pixel[10]+offset]) > cb {
																					if int(im[pixel[11]+offset]) > cb {
																						if int(im[pixel[12]+offset]) > cb {
																							if int(im[pixel[14]+offset]) > cb {
																								goto is_a_corner
																							} else {
																								goto is_not_a_corner
																							}
																						} else {
																							goto is_not_a_corner
																						}
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	if int(im[pixel[13]+offset]) < c_b {
																		if int(im[pixel[14]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																}
															} else {
																if int(im[pixel[7]+offset]) > cb {
																	if int(im[pixel[8]+offset]) > cb {
																		if int(im[pixel[9]+offset]) > cb {
																			if int(im[pixel[10]+offset]) > cb {
																				if int(im[pixel[11]+offset]) > cb {
																					if int(im[pixel[12]+offset]) > cb {
																						if int(im[pixel[13]+offset]) > cb {
																							if int(im[pixel[14]+offset]) > cb {
																								goto is_a_corner
																							} else {
																								goto is_not_a_corner
																							}
																						} else {
																							goto is_not_a_corner
																						}
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															}
														} else {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) > cb {
																	if int(im[pixel[14]+offset]) > cb {
																		if int(im[pixel[8]+offset]) > cb {
																			if int(im[pixel[9]+offset]) > cb {
																				if int(im[pixel[10]+offset]) > cb {
																					if int(im[pixel[11]+offset]) > cb {
																						if int(im[pixel[12]+offset]) > cb {
																							if int(im[pixel[13]+offset]) > cb {
																								if int(im[pixel[15]+offset]) > cb {
																									goto is_a_corner
																								} else {
																									goto is_not_a_corner
																								}
																							} else {
																								goto is_not_a_corner
																							}
																						} else {
																							goto is_not_a_corner
																						}
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		if int(im[pixel[14]+offset]) < c_b {
																			if int(im[pixel[15]+offset]) < c_b {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	}
																} else {
																	if int(im[pixel[7]+offset]) < c_b {
																		if int(im[pixel[8]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			if int(im[pixel[15]+offset]) < c_b {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		}
																	} else {
																		if int(im[pixel[14]+offset]) < c_b {
																			if int(im[pixel[15]+offset]) < c_b {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	}
																}
															} else {
																if int(im[pixel[13]+offset]) > cb {
																	if int(im[pixel[7]+offset]) > cb {
																		if int(im[pixel[8]+offset]) > cb {
																			if int(im[pixel[9]+offset]) > cb {
																				if int(im[pixel[10]+offset]) > cb {
																					if int(im[pixel[11]+offset]) > cb {
																						if int(im[pixel[12]+offset]) > cb {
																							if int(im[pixel[14]+offset]) > cb {
																								if int(im[pixel[15]+offset]) > cb {
																									goto is_a_corner
																								} else {
																									goto is_not_a_corner
																								}
																							} else {
																								goto is_not_a_corner
																							}
																						} else {
																							goto is_not_a_corner
																						}
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	if int(im[pixel[13]+offset]) < c_b {
																		if int(im[pixel[14]+offset]) < c_b {
																			if int(im[pixel[15]+offset]) < c_b {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																}
															}
														}
													} else {
														if int(im[pixel[12]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																		if int(im[pixel[10]+offset]) > cb {
																			if int(im[pixel[11]+offset]) > cb {
																				if int(im[pixel[13]+offset]) > cb {
																					if int(im[pixel[14]+offset]) > cb {
																						if int(im[pixel[6]+offset]) > cb {
																							goto is_a_corner
																						} else {
																							if int(im[pixel[15]+offset]) > cb {
																								goto is_a_corner
																							} else {
																								goto is_not_a_corner
																							}
																						}
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															if int(im[pixel[12]+offset]) < c_b {
																if int(im[pixel[13]+offset]) < c_b {
																	if int(im[pixel[14]+offset]) < c_b {
																		if int(im[pixel[15]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			if int(im[pixel[6]+offset]) < c_b {
																				if int(im[pixel[7]+offset]) < c_b {
																					if int(im[pixel[8]+offset]) < c_b {
																						if int(im[pixel[9]+offset]) < c_b {
																							if int(im[pixel[10]+offset]) < c_b {
																								if int(im[pixel[11]+offset]) < c_b {
																									goto is_a_corner
																								} else {
																									goto is_not_a_corner
																								}
																							} else {
																								goto is_not_a_corner
																							}
																						} else {
																							goto is_not_a_corner
																						}
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													}
												}
											} else {
												if int(im[pixel[11]+offset]) > cb {
													if int(im[pixel[7]+offset]) > cb {
														if int(im[pixel[8]+offset]) > cb {
															if int(im[pixel[9]+offset]) > cb {
																if int(im[pixel[10]+offset]) > cb {
																	if int(im[pixel[12]+offset]) > cb {
																		if int(im[pixel[13]+offset]) > cb {
																			if int(im[pixel[6]+offset]) > cb {
																				if int(im[pixel[5]+offset]) > cb {
																					goto is_a_corner
																				} else {
																					if int(im[pixel[14]+offset]) > cb {
																						goto is_a_corner
																					} else {
																						goto is_not_a_corner
																					}
																				}
																			} else {
																				if int(im[pixel[14]+offset]) > cb {
																					if int(im[pixel[15]+offset]) > cb {
																						goto is_a_corner
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													if int(im[pixel[11]+offset]) < c_b {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[13]+offset]) < c_b {
																if int(im[pixel[14]+offset]) < c_b {
																	if int(im[pixel[15]+offset]) < c_b {
																		goto is_a_corner
																	} else {
																		if int(im[pixel[6]+offset]) < c_b {
																			if int(im[pixel[7]+offset]) < c_b {
																				if int(im[pixel[8]+offset]) < c_b {
																					if int(im[pixel[9]+offset]) < c_b {
																						if int(im[pixel[10]+offset]) < c_b {
																							goto is_a_corner
																						} else {
																							goto is_not_a_corner
																						}
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	}
																} else {
																	if int(im[pixel[5]+offset]) < c_b {
																		if int(im[pixel[6]+offset]) < c_b {
																			if int(im[pixel[7]+offset]) < c_b {
																				if int(im[pixel[8]+offset]) < c_b {
																					if int(im[pixel[9]+offset]) < c_b {
																						if int(im[pixel[10]+offset]) < c_b {
																							goto is_a_corner
																						} else {
																							goto is_not_a_corner
																						}
																					} else {
																						goto is_not_a_corner
																					}
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											}
										}
									} else {
										if int(im[pixel[10]+offset]) > cb {
											if int(im[pixel[7]+offset]) > cb {
												if int(im[pixel[8]+offset]) > cb {
													if int(im[pixel[9]+offset]) > cb {
														if int(im[pixel[11]+offset]) > cb {
															if int(im[pixel[12]+offset]) > cb {
																if int(im[pixel[6]+offset]) > cb {
																	if int(im[pixel[5]+offset]) > cb {
																		if int(im[pixel[4]+offset]) > cb {
																			goto is_a_corner
																		} else {
																			if int(im[pixel[13]+offset]) > cb {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		}
																	} else {
																		if int(im[pixel[13]+offset]) > cb {
																			if int(im[pixel[14]+offset]) > cb {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	}
																} else {
																	if int(im[pixel[13]+offset]) > cb {
																		if int(im[pixel[14]+offset]) > cb {
																			if int(im[pixel[15]+offset]) > cb {
																				goto is_a_corner
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											if int(im[pixel[10]+offset]) < c_b {
												if int(im[pixel[11]+offset]) < c_b {
													if int(im[pixel[12]+offset]) < c_b {
														if int(im[pixel[13]+offset]) < c_b {
															if int(im[pixel[14]+offset]) < c_b {
																if int(im[pixel[15]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	if int(im[pixel[6]+offset]) < c_b {
																		if int(im[pixel[7]+offset]) < c_b {
																			if int(im[pixel[8]+offset]) < c_b {
																				if int(im[pixel[9]+offset]) < c_b {
																					goto is_a_corner
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																}
															} else {
																if int(im[pixel[5]+offset]) < c_b {
																	if int(im[pixel[6]+offset]) < c_b {
																		if int(im[pixel[7]+offset]) < c_b {
																			if int(im[pixel[8]+offset]) < c_b {
																				if int(im[pixel[9]+offset]) < c_b {
																					goto is_a_corner
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															}
														} else {
															if int(im[pixel[4]+offset]) < c_b {
																if int(im[pixel[5]+offset]) < c_b {
																	if int(im[pixel[6]+offset]) < c_b {
																		if int(im[pixel[7]+offset]) < c_b {
																			if int(im[pixel[8]+offset]) < c_b {
																				if int(im[pixel[9]+offset]) < c_b {
																					goto is_a_corner
																				} else {
																					goto is_not_a_corner
																				}
																			} else {
																				goto is_not_a_corner
																			}
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									}
								}
							} else {
								if int(im[pixel[9]+offset]) > cb {
									if int(im[pixel[7]+offset]) > cb {
										if int(im[pixel[8]+offset]) > cb {
											if int(im[pixel[10]+offset]) > cb {
												if int(im[pixel[11]+offset]) > cb {
													if int(im[pixel[6]+offset]) > cb {
														if int(im[pixel[5]+offset]) > cb {
															if int(im[pixel[4]+offset]) > cb {
																if int(im[pixel[3]+offset]) > cb {
																	goto is_a_corner
																} else {
																	if int(im[pixel[12]+offset]) > cb {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																}
															} else {
																if int(im[pixel[12]+offset]) > cb {
																	if int(im[pixel[13]+offset]) > cb {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															}
														} else {
															if int(im[pixel[12]+offset]) > cb {
																if int(im[pixel[13]+offset]) > cb {
																	if int(im[pixel[14]+offset]) > cb {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														if int(im[pixel[12]+offset]) > cb {
															if int(im[pixel[13]+offset]) > cb {
																if int(im[pixel[14]+offset]) > cb {
																	if int(im[pixel[15]+offset]) > cb {
																		goto is_a_corner
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									if int(im[pixel[9]+offset]) < c_b {
										if int(im[pixel[10]+offset]) < c_b {
											if int(im[pixel[11]+offset]) < c_b {
												if int(im[pixel[12]+offset]) < c_b {
													if int(im[pixel[13]+offset]) < c_b {
														if int(im[pixel[14]+offset]) < c_b {
															if int(im[pixel[15]+offset]) < c_b {
																goto is_a_corner
															} else {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																		if int(im[pixel[8]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															}
														} else {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																		if int(im[pixel[8]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														if int(im[pixel[4]+offset]) < c_b {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																		if int(im[pixel[8]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													if int(im[pixel[3]+offset]) < c_b {
														if int(im[pixel[4]+offset]) < c_b {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																		if int(im[pixel[8]+offset]) < c_b {
																			goto is_a_corner
																		} else {
																			goto is_not_a_corner
																		}
																	} else {
																		goto is_not_a_corner
																	}
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								}
							}
						}
					} else {
						if int(im[pixel[8]+offset]) > cb {
							if int(im[pixel[7]+offset]) > cb {
								if int(im[pixel[9]+offset]) > cb {
									if int(im[pixel[10]+offset]) > cb {
										if int(im[pixel[6]+offset]) > cb {
											if int(im[pixel[5]+offset]) > cb {
												if int(im[pixel[4]+offset]) > cb {
													if int(im[pixel[3]+offset]) > cb {
														if int(im[pixel[2]+offset]) > cb {
															goto is_a_corner
														} else {
															if int(im[pixel[11]+offset]) > cb {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														if int(im[pixel[11]+offset]) > cb {
															if int(im[pixel[12]+offset]) > cb {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													if int(im[pixel[11]+offset]) > cb {
														if int(im[pixel[12]+offset]) > cb {
															if int(im[pixel[13]+offset]) > cb {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[11]+offset]) > cb {
													if int(im[pixel[12]+offset]) > cb {
														if int(im[pixel[13]+offset]) > cb {
															if int(im[pixel[14]+offset]) > cb {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[11]+offset]) > cb {
												if int(im[pixel[12]+offset]) > cb {
													if int(im[pixel[13]+offset]) > cb {
														if int(im[pixel[14]+offset]) > cb {
															if int(im[pixel[15]+offset]) > cb {
																goto is_a_corner
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									goto is_not_a_corner
								}
							} else {
								goto is_not_a_corner
							}
						} else {
							if int(im[pixel[8]+offset]) < c_b {
								if int(im[pixel[9]+offset]) < c_b {
									if int(im[pixel[10]+offset]) < c_b {
										if int(im[pixel[11]+offset]) < c_b {
											if int(im[pixel[12]+offset]) < c_b {
												if int(im[pixel[13]+offset]) < c_b {
													if int(im[pixel[14]+offset]) < c_b {
														if int(im[pixel[15]+offset]) < c_b {
															goto is_a_corner
														} else {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														}
													} else {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													if int(im[pixel[4]+offset]) < c_b {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[3]+offset]) < c_b {
													if int(im[pixel[4]+offset]) < c_b {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[2]+offset]) < c_b {
												if int(im[pixel[3]+offset]) < c_b {
													if int(im[pixel[4]+offset]) < c_b {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																	goto is_a_corner
																} else {
																	goto is_not_a_corner
																}
															} else {
																goto is_not_a_corner
															}
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									goto is_not_a_corner
								}
							} else {
								goto is_not_a_corner
							}
						}
					}
				}
			} else {
				if int(im[pixel[7]+offset]) > cb {
					if int(im[pixel[8]+offset]) > cb {
						if int(im[pixel[9]+offset]) > cb {
							if int(im[pixel[6]+offset]) > cb {
								if int(im[pixel[5]+offset]) > cb {
									if int(im[pixel[4]+offset]) > cb {
										if int(im[pixel[3]+offset]) > cb {
											if int(im[pixel[2]+offset]) > cb {
												if int(im[pixel[1]+offset]) > cb {
													goto is_a_corner
												} else {
													if int(im[pixel[10]+offset]) > cb {
														goto is_a_corner
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[10]+offset]) > cb {
													if int(im[pixel[11]+offset]) > cb {
														goto is_a_corner
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[10]+offset]) > cb {
												if int(im[pixel[11]+offset]) > cb {
													if int(im[pixel[12]+offset]) > cb {
														goto is_a_corner
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										if int(im[pixel[10]+offset]) > cb {
											if int(im[pixel[11]+offset]) > cb {
												if int(im[pixel[12]+offset]) > cb {
													if int(im[pixel[13]+offset]) > cb {
														goto is_a_corner
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									}
								} else {
									if int(im[pixel[10]+offset]) > cb {
										if int(im[pixel[11]+offset]) > cb {
											if int(im[pixel[12]+offset]) > cb {
												if int(im[pixel[13]+offset]) > cb {
													if int(im[pixel[14]+offset]) > cb {
														goto is_a_corner
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								}
							} else {
								if int(im[pixel[10]+offset]) > cb {
									if int(im[pixel[11]+offset]) > cb {
										if int(im[pixel[12]+offset]) > cb {
											if int(im[pixel[13]+offset]) > cb {
												if int(im[pixel[14]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
														goto is_a_corner
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								} else {
									goto is_not_a_corner
								}
							}
						} else {
							goto is_not_a_corner
						}
					} else {
						goto is_not_a_corner
					}
				} else {
					if int(im[pixel[7]+offset]) < c_b {
						if int(im[pixel[8]+offset]) < c_b {
							if int(im[pixel[9]+offset]) < c_b {
								if int(im[pixel[6]+offset]) < c_b {
									if int(im[pixel[5]+offset]) < c_b {
										if int(im[pixel[4]+offset]) < c_b {
											if int(im[pixel[3]+offset]) < c_b {
												if int(im[pixel[2]+offset]) < c_b {
													if int(im[pixel[1]+offset]) < c_b {
														goto is_a_corner
													} else {
														if int(im[pixel[10]+offset]) < c_b {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													}
												} else {
													if int(im[pixel[10]+offset]) < c_b {
														if int(im[pixel[11]+offset]) < c_b {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												}
											} else {
												if int(im[pixel[10]+offset]) < c_b {
													if int(im[pixel[11]+offset]) < c_b {
														if int(im[pixel[12]+offset]) < c_b {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											}
										} else {
											if int(im[pixel[10]+offset]) < c_b {
												if int(im[pixel[11]+offset]) < c_b {
													if int(im[pixel[12]+offset]) < c_b {
														if int(im[pixel[13]+offset]) < c_b {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										}
									} else {
										if int(im[pixel[10]+offset]) < c_b {
											if int(im[pixel[11]+offset]) < c_b {
												if int(im[pixel[12]+offset]) < c_b {
													if int(im[pixel[13]+offset]) < c_b {
														if int(im[pixel[14]+offset]) < c_b {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									}
								} else {
									if int(im[pixel[10]+offset]) < c_b {
										if int(im[pixel[11]+offset]) < c_b {
											if int(im[pixel[12]+offset]) < c_b {
												if int(im[pixel[13]+offset]) < c_b {
													if int(im[pixel[14]+offset]) < c_b {
														if int(im[pixel[15]+offset]) < c_b {
															goto is_a_corner
														} else {
															goto is_not_a_corner
														}
													} else {
														goto is_not_a_corner
													}
												} else {
													goto is_not_a_corner
												}
											} else {
												goto is_not_a_corner
											}
										} else {
											goto is_not_a_corner
										}
									} else {
										goto is_not_a_corner
									}
								}
							} else {
								goto is_not_a_corner
							}
						} else {
							goto is_not_a_corner
						}
					} else {
						goto is_not_a_corner
					}
				}
			}
		}
	is_a_corner:
		;
		bmin = b
		goto end_if
	is_not_a_corner:
		;
		bmax = b
		goto end_if
	end_if:
		;
		if bmin == bmax-1 || bmin == bmax {
			return int(bmin)
		}
		b = (bmin + bmax) / 2
	}
	return
}

// make_offsets - transpiled function from  fast_9.c:2938
func make_offsets(pixel []int, row_stride int) {
	pixel[0] = 0 + row_stride*3
	pixel[1] = 1 + row_stride*3
	pixel[2] = 2 + row_stride*2
	pixel[3] = 3 + row_stride*1
	pixel[4] = 3 + row_stride*0
	pixel[5] = 3 + row_stride*-1
	pixel[6] = 2 + row_stride*-2
	pixel[7] = 1 + row_stride*-3
	pixel[8] = 0 + row_stride*-3
	pixel[9] = -1 + row_stride*-3
	pixel[10] = -2 + row_stride*-2
	pixel[11] = -3 + row_stride*-1
	pixel[12] = -3 + row_stride*0
	pixel[13] = -3 + row_stride*1
	pixel[14] = -2 + row_stride*2
	pixel[15] = -1 + row_stride*3
}

// fast9_score - transpiled function from  fast_9.c:2960
func fast9_score(i []uint8, stride int, corners []image.Point, b int) []int {
	var scores []int = make([]int, len(corners))
	var pixel []int = make([]int, 16)
	make_offsets(pixel, stride)
	for n := 0; n < len(corners); n++ {
		scores[n] = fast9_corner_score(i, corners[n].Y*stride+corners[n].X, pixel, b)
	}
	return scores
}

// fast9_detect - transpiled function from  fast_9.c:2975
func fast9_detect(im []uint8, xsize, ysize, stride, b int) []image.Point {
	var rsize int = 512
	var pixel []int = make([]int, 16, 16)
	ret_corners := make([]image.Point, 0, rsize)
	make_offsets(pixel, stride)
	x, y := 0, 0
	for y = 3; y < ysize-3; y++ {
		for x = 3; x < xsize-3; x++ {
			offset := y*stride + x
			var cb int = int(im[offset]) + b
			var c_b int = int(im[offset]) - b
			if int(im[pixel[0]+offset]) > cb {
				if int(im[pixel[1]+offset]) > cb {
					if int(im[pixel[2]+offset]) > cb {
						if int(im[pixel[3]+offset]) > cb {
							if int(im[pixel[4]+offset]) > cb {
								if int(im[pixel[5]+offset]) > cb {
									if int(im[pixel[6]+offset]) > cb {
										if int(im[pixel[7]+offset]) > cb {
											if int(im[pixel[8]+offset]) > cb {
											} else {
												if int(im[pixel[15]+offset]) > cb {
												} else {
													continue
												}
											}
										} else {
											if int(im[pixel[7]+offset]) < c_b {
												if int(im[pixel[14]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
													} else {
														continue
													}
												} else {
													if int(im[pixel[14]+offset]) < c_b {
														if int(im[pixel[8]+offset]) < c_b {
															if int(im[pixel[9]+offset]) < c_b {
																if int(im[pixel[10]+offset]) < c_b {
																	if int(im[pixel[11]+offset]) < c_b {
																		if int(im[pixel[12]+offset]) < c_b {
																			if int(im[pixel[13]+offset]) < c_b {
																				if int(im[pixel[15]+offset]) < c_b {
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[14]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
													} else {
														continue
													}
												} else {
													continue
												}
											}
										}
									} else {
										if int(im[pixel[6]+offset]) < c_b {
											if int(im[pixel[15]+offset]) > cb {
												if int(im[pixel[13]+offset]) > cb {
													if int(im[pixel[14]+offset]) > cb {
													} else {
														continue
													}
												} else {
													if int(im[pixel[13]+offset]) < c_b {
														if int(im[pixel[7]+offset]) < c_b {
															if int(im[pixel[8]+offset]) < c_b {
																if int(im[pixel[9]+offset]) < c_b {
																	if int(im[pixel[10]+offset]) < c_b {
																		if int(im[pixel[11]+offset]) < c_b {
																			if int(im[pixel[12]+offset]) < c_b {
																				if int(im[pixel[14]+offset]) < c_b {
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[7]+offset]) < c_b {
													if int(im[pixel[8]+offset]) < c_b {
														if int(im[pixel[9]+offset]) < c_b {
															if int(im[pixel[10]+offset]) < c_b {
																if int(im[pixel[11]+offset]) < c_b {
																	if int(im[pixel[12]+offset]) < c_b {
																		if int(im[pixel[13]+offset]) < c_b {
																			if int(im[pixel[14]+offset]) < c_b {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											if int(im[pixel[13]+offset]) > cb {
												if int(im[pixel[14]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												if int(im[pixel[13]+offset]) < c_b {
													if int(im[pixel[7]+offset]) < c_b {
														if int(im[pixel[8]+offset]) < c_b {
															if int(im[pixel[9]+offset]) < c_b {
																if int(im[pixel[10]+offset]) < c_b {
																	if int(im[pixel[11]+offset]) < c_b {
																		if int(im[pixel[12]+offset]) < c_b {
																			if int(im[pixel[14]+offset]) < c_b {
																				if int(im[pixel[15]+offset]) < c_b {
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										}
									}
								} else {
									if int(im[pixel[5]+offset]) < c_b {
										if int(im[pixel[14]+offset]) > cb {
											if int(im[pixel[12]+offset]) > cb {
												if int(im[pixel[13]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
													} else {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																		if int(im[pixel[10]+offset]) > cb {
																			if int(im[pixel[11]+offset]) > cb {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													continue
												}
											} else {
												if int(im[pixel[12]+offset]) < c_b {
													if int(im[pixel[6]+offset]) < c_b {
														if int(im[pixel[7]+offset]) < c_b {
															if int(im[pixel[8]+offset]) < c_b {
																if int(im[pixel[9]+offset]) < c_b {
																	if int(im[pixel[10]+offset]) < c_b {
																		if int(im[pixel[11]+offset]) < c_b {
																			if int(im[pixel[13]+offset]) < c_b {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											if int(im[pixel[14]+offset]) < c_b {
												if int(im[pixel[7]+offset]) < c_b {
													if int(im[pixel[8]+offset]) < c_b {
														if int(im[pixel[9]+offset]) < c_b {
															if int(im[pixel[10]+offset]) < c_b {
																if int(im[pixel[11]+offset]) < c_b {
																	if int(im[pixel[12]+offset]) < c_b {
																		if int(im[pixel[13]+offset]) < c_b {
																			if int(im[pixel[6]+offset]) < c_b {
																			} else {
																				if int(im[pixel[15]+offset]) < c_b {
																				} else {
																					continue
																				}
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												if int(im[pixel[6]+offset]) < c_b {
													if int(im[pixel[7]+offset]) < c_b {
														if int(im[pixel[8]+offset]) < c_b {
															if int(im[pixel[9]+offset]) < c_b {
																if int(im[pixel[10]+offset]) < c_b {
																	if int(im[pixel[11]+offset]) < c_b {
																		if int(im[pixel[12]+offset]) < c_b {
																			if int(im[pixel[13]+offset]) < c_b {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										}
									} else {
										if int(im[pixel[12]+offset]) > cb {
											if int(im[pixel[13]+offset]) > cb {
												if int(im[pixel[14]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
													} else {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																		if int(im[pixel[10]+offset]) > cb {
																			if int(im[pixel[11]+offset]) > cb {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											if int(im[pixel[12]+offset]) < c_b {
												if int(im[pixel[7]+offset]) < c_b {
													if int(im[pixel[8]+offset]) < c_b {
														if int(im[pixel[9]+offset]) < c_b {
															if int(im[pixel[10]+offset]) < c_b {
																if int(im[pixel[11]+offset]) < c_b {
																	if int(im[pixel[13]+offset]) < c_b {
																		if int(im[pixel[14]+offset]) < c_b {
																			if int(im[pixel[6]+offset]) < c_b {
																			} else {
																				if int(im[pixel[15]+offset]) < c_b {
																				} else {
																					continue
																				}
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										}
									}
								}
							} else {
								if int(im[pixel[4]+offset]) < c_b {
									if int(im[pixel[13]+offset]) > cb {
										if int(im[pixel[11]+offset]) > cb {
											if int(im[pixel[12]+offset]) > cb {
												if int(im[pixel[14]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
													} else {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																		if int(im[pixel[10]+offset]) > cb {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																		if int(im[pixel[10]+offset]) > cb {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												continue
											}
										} else {
											if int(im[pixel[11]+offset]) < c_b {
												if int(im[pixel[5]+offset]) < c_b {
													if int(im[pixel[6]+offset]) < c_b {
														if int(im[pixel[7]+offset]) < c_b {
															if int(im[pixel[8]+offset]) < c_b {
																if int(im[pixel[9]+offset]) < c_b {
																	if int(im[pixel[10]+offset]) < c_b {
																		if int(im[pixel[12]+offset]) < c_b {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										}
									} else {
										if int(im[pixel[13]+offset]) < c_b {
											if int(im[pixel[7]+offset]) < c_b {
												if int(im[pixel[8]+offset]) < c_b {
													if int(im[pixel[9]+offset]) < c_b {
														if int(im[pixel[10]+offset]) < c_b {
															if int(im[pixel[11]+offset]) < c_b {
																if int(im[pixel[12]+offset]) < c_b {
																	if int(im[pixel[6]+offset]) < c_b {
																		if int(im[pixel[5]+offset]) < c_b {
																		} else {
																			if int(im[pixel[14]+offset]) < c_b {
																			} else {
																				continue
																			}
																		}
																	} else {
																		if int(im[pixel[14]+offset]) < c_b {
																			if int(im[pixel[15]+offset]) < c_b {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											if int(im[pixel[5]+offset]) < c_b {
												if int(im[pixel[6]+offset]) < c_b {
													if int(im[pixel[7]+offset]) < c_b {
														if int(im[pixel[8]+offset]) < c_b {
															if int(im[pixel[9]+offset]) < c_b {
																if int(im[pixel[10]+offset]) < c_b {
																	if int(im[pixel[11]+offset]) < c_b {
																		if int(im[pixel[12]+offset]) < c_b {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										}
									}
								} else {
									if int(im[pixel[11]+offset]) > cb {
										if int(im[pixel[12]+offset]) > cb {
											if int(im[pixel[13]+offset]) > cb {
												if int(im[pixel[14]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
													} else {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																		if int(im[pixel[10]+offset]) > cb {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																		if int(im[pixel[10]+offset]) > cb {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												continue
											}
										} else {
											continue
										}
									} else {
										if int(im[pixel[11]+offset]) < c_b {
											if int(im[pixel[7]+offset]) < c_b {
												if int(im[pixel[8]+offset]) < c_b {
													if int(im[pixel[9]+offset]) < c_b {
														if int(im[pixel[10]+offset]) < c_b {
															if int(im[pixel[12]+offset]) < c_b {
																if int(im[pixel[13]+offset]) < c_b {
																	if int(im[pixel[6]+offset]) < c_b {
																		if int(im[pixel[5]+offset]) < c_b {
																		} else {
																			if int(im[pixel[14]+offset]) < c_b {
																			} else {
																				continue
																			}
																		}
																	} else {
																		if int(im[pixel[14]+offset]) < c_b {
																			if int(im[pixel[15]+offset]) < c_b {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											continue
										}
									}
								}
							}
						} else {
							if int(im[pixel[3]+offset]) < c_b {
								if int(im[pixel[10]+offset]) > cb {
									if int(im[pixel[11]+offset]) > cb {
										if int(im[pixel[12]+offset]) > cb {
											if int(im[pixel[13]+offset]) > cb {
												if int(im[pixel[14]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
													} else {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[4]+offset]) > cb {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											continue
										}
									} else {
										continue
									}
								} else {
									if int(im[pixel[10]+offset]) < c_b {
										if int(im[pixel[7]+offset]) < c_b {
											if int(im[pixel[8]+offset]) < c_b {
												if int(im[pixel[9]+offset]) < c_b {
													if int(im[pixel[11]+offset]) < c_b {
														if int(im[pixel[6]+offset]) < c_b {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[4]+offset]) < c_b {
																} else {
																	if int(im[pixel[12]+offset]) < c_b {
																		if int(im[pixel[13]+offset]) < c_b {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																}
															} else {
																if int(im[pixel[12]+offset]) < c_b {
																	if int(im[pixel[13]+offset]) < c_b {
																		if int(im[pixel[14]+offset]) < c_b {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														} else {
															if int(im[pixel[12]+offset]) < c_b {
																if int(im[pixel[13]+offset]) < c_b {
																	if int(im[pixel[14]+offset]) < c_b {
																		if int(im[pixel[15]+offset]) < c_b {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											continue
										}
									} else {
										continue
									}
								}
							} else {
								if int(im[pixel[10]+offset]) > cb {
									if int(im[pixel[11]+offset]) > cb {
										if int(im[pixel[12]+offset]) > cb {
											if int(im[pixel[13]+offset]) > cb {
												if int(im[pixel[14]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
													} else {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[4]+offset]) > cb {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											continue
										}
									} else {
										continue
									}
								} else {
									if int(im[pixel[10]+offset]) < c_b {
										if int(im[pixel[7]+offset]) < c_b {
											if int(im[pixel[8]+offset]) < c_b {
												if int(im[pixel[9]+offset]) < c_b {
													if int(im[pixel[11]+offset]) < c_b {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[5]+offset]) < c_b {
																	if int(im[pixel[4]+offset]) < c_b {
																	} else {
																		if int(im[pixel[13]+offset]) < c_b {
																		} else {
																			continue
																		}
																	}
																} else {
																	if int(im[pixel[13]+offset]) < c_b {
																		if int(im[pixel[14]+offset]) < c_b {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																}
															} else {
																if int(im[pixel[13]+offset]) < c_b {
																	if int(im[pixel[14]+offset]) < c_b {
																		if int(im[pixel[15]+offset]) < c_b {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											continue
										}
									} else {
										continue
									}
								}
							}
						}
					} else {
						if int(im[pixel[2]+offset]) < c_b {
							if int(im[pixel[9]+offset]) > cb {
								if int(im[pixel[10]+offset]) > cb {
									if int(im[pixel[11]+offset]) > cb {
										if int(im[pixel[12]+offset]) > cb {
											if int(im[pixel[13]+offset]) > cb {
												if int(im[pixel[14]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
													} else {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[4]+offset]) > cb {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											if int(im[pixel[3]+offset]) > cb {
												if int(im[pixel[4]+offset]) > cb {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										}
									} else {
										continue
									}
								} else {
									continue
								}
							} else {
								if int(im[pixel[9]+offset]) < c_b {
									if int(im[pixel[7]+offset]) < c_b {
										if int(im[pixel[8]+offset]) < c_b {
											if int(im[pixel[10]+offset]) < c_b {
												if int(im[pixel[6]+offset]) < c_b {
													if int(im[pixel[5]+offset]) < c_b {
														if int(im[pixel[4]+offset]) < c_b {
															if int(im[pixel[3]+offset]) < c_b {
															} else {
																if int(im[pixel[11]+offset]) < c_b {
																	if int(im[pixel[12]+offset]) < c_b {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														} else {
															if int(im[pixel[11]+offset]) < c_b {
																if int(im[pixel[12]+offset]) < c_b {
																	if int(im[pixel[13]+offset]) < c_b {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														if int(im[pixel[11]+offset]) < c_b {
															if int(im[pixel[12]+offset]) < c_b {
																if int(im[pixel[13]+offset]) < c_b {
																	if int(im[pixel[14]+offset]) < c_b {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[11]+offset]) < c_b {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[13]+offset]) < c_b {
																if int(im[pixel[14]+offset]) < c_b {
																	if int(im[pixel[15]+offset]) < c_b {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												continue
											}
										} else {
											continue
										}
									} else {
										continue
									}
								} else {
									continue
								}
							}
						} else {
							if int(im[pixel[9]+offset]) > cb {
								if int(im[pixel[10]+offset]) > cb {
									if int(im[pixel[11]+offset]) > cb {
										if int(im[pixel[12]+offset]) > cb {
											if int(im[pixel[13]+offset]) > cb {
												if int(im[pixel[14]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
													} else {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[4]+offset]) > cb {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											if int(im[pixel[3]+offset]) > cb {
												if int(im[pixel[4]+offset]) > cb {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										}
									} else {
										continue
									}
								} else {
									continue
								}
							} else {
								if int(im[pixel[9]+offset]) < c_b {
									if int(im[pixel[7]+offset]) < c_b {
										if int(im[pixel[8]+offset]) < c_b {
											if int(im[pixel[10]+offset]) < c_b {
												if int(im[pixel[11]+offset]) < c_b {
													if int(im[pixel[6]+offset]) < c_b {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[4]+offset]) < c_b {
																if int(im[pixel[3]+offset]) < c_b {
																} else {
																	if int(im[pixel[12]+offset]) < c_b {
																	} else {
																		continue
																	}
																}
															} else {
																if int(im[pixel[12]+offset]) < c_b {
																	if int(im[pixel[13]+offset]) < c_b {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														} else {
															if int(im[pixel[12]+offset]) < c_b {
																if int(im[pixel[13]+offset]) < c_b {
																	if int(im[pixel[14]+offset]) < c_b {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[13]+offset]) < c_b {
																if int(im[pixel[14]+offset]) < c_b {
																	if int(im[pixel[15]+offset]) < c_b {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											continue
										}
									} else {
										continue
									}
								} else {
									continue
								}
							}
						}
					}
				} else {
					if int(im[pixel[1]+offset]) < c_b {
						if int(im[pixel[8]+offset]) > cb {
							if int(im[pixel[9]+offset]) > cb {
								if int(im[pixel[10]+offset]) > cb {
									if int(im[pixel[11]+offset]) > cb {
										if int(im[pixel[12]+offset]) > cb {
											if int(im[pixel[13]+offset]) > cb {
												if int(im[pixel[14]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
													} else {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[4]+offset]) > cb {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											if int(im[pixel[3]+offset]) > cb {
												if int(im[pixel[4]+offset]) > cb {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										}
									} else {
										if int(im[pixel[2]+offset]) > cb {
											if int(im[pixel[3]+offset]) > cb {
												if int(im[pixel[4]+offset]) > cb {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											continue
										}
									}
								} else {
									continue
								}
							} else {
								continue
							}
						} else {
							if int(im[pixel[8]+offset]) < c_b {
								if int(im[pixel[7]+offset]) < c_b {
									if int(im[pixel[9]+offset]) < c_b {
										if int(im[pixel[6]+offset]) < c_b {
											if int(im[pixel[5]+offset]) < c_b {
												if int(im[pixel[4]+offset]) < c_b {
													if int(im[pixel[3]+offset]) < c_b {
														if int(im[pixel[2]+offset]) < c_b {
														} else {
															if int(im[pixel[10]+offset]) < c_b {
																if int(im[pixel[11]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														if int(im[pixel[10]+offset]) < c_b {
															if int(im[pixel[11]+offset]) < c_b {
																if int(im[pixel[12]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[10]+offset]) < c_b {
														if int(im[pixel[11]+offset]) < c_b {
															if int(im[pixel[12]+offset]) < c_b {
																if int(im[pixel[13]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[10]+offset]) < c_b {
													if int(im[pixel[11]+offset]) < c_b {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[13]+offset]) < c_b {
																if int(im[pixel[14]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											if int(im[pixel[10]+offset]) < c_b {
												if int(im[pixel[11]+offset]) < c_b {
													if int(im[pixel[12]+offset]) < c_b {
														if int(im[pixel[13]+offset]) < c_b {
															if int(im[pixel[14]+offset]) < c_b {
																if int(im[pixel[15]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										}
									} else {
										continue
									}
								} else {
									continue
								}
							} else {
								continue
							}
						}
					} else {
						if int(im[pixel[8]+offset]) > cb {
							if int(im[pixel[9]+offset]) > cb {
								if int(im[pixel[10]+offset]) > cb {
									if int(im[pixel[11]+offset]) > cb {
										if int(im[pixel[12]+offset]) > cb {
											if int(im[pixel[13]+offset]) > cb {
												if int(im[pixel[14]+offset]) > cb {
													if int(im[pixel[15]+offset]) > cb {
													} else {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[4]+offset]) > cb {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											if int(im[pixel[3]+offset]) > cb {
												if int(im[pixel[4]+offset]) > cb {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										}
									} else {
										if int(im[pixel[2]+offset]) > cb {
											if int(im[pixel[3]+offset]) > cb {
												if int(im[pixel[4]+offset]) > cb {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											continue
										}
									}
								} else {
									continue
								}
							} else {
								continue
							}
						} else {
							if int(im[pixel[8]+offset]) < c_b {
								if int(im[pixel[7]+offset]) < c_b {
									if int(im[pixel[9]+offset]) < c_b {
										if int(im[pixel[10]+offset]) < c_b {
											if int(im[pixel[6]+offset]) < c_b {
												if int(im[pixel[5]+offset]) < c_b {
													if int(im[pixel[4]+offset]) < c_b {
														if int(im[pixel[3]+offset]) < c_b {
															if int(im[pixel[2]+offset]) < c_b {
															} else {
																if int(im[pixel[11]+offset]) < c_b {
																} else {
																	continue
																}
															}
														} else {
															if int(im[pixel[11]+offset]) < c_b {
																if int(im[pixel[12]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														if int(im[pixel[11]+offset]) < c_b {
															if int(im[pixel[12]+offset]) < c_b {
																if int(im[pixel[13]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[11]+offset]) < c_b {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[13]+offset]) < c_b {
																if int(im[pixel[14]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[11]+offset]) < c_b {
													if int(im[pixel[12]+offset]) < c_b {
														if int(im[pixel[13]+offset]) < c_b {
															if int(im[pixel[14]+offset]) < c_b {
																if int(im[pixel[15]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											continue
										}
									} else {
										continue
									}
								} else {
									continue
								}
							} else {
								continue
							}
						}
					}
				}
			} else {
				if int(im[pixel[0]+offset]) < c_b {
					if int(im[pixel[1]+offset]) > cb {
						if int(im[pixel[8]+offset]) > cb {
							if int(im[pixel[7]+offset]) > cb {
								if int(im[pixel[9]+offset]) > cb {
									if int(im[pixel[6]+offset]) > cb {
										if int(im[pixel[5]+offset]) > cb {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[3]+offset]) > cb {
													if int(im[pixel[2]+offset]) > cb {
													} else {
														if int(im[pixel[10]+offset]) > cb {
															if int(im[pixel[11]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[10]+offset]) > cb {
														if int(im[pixel[11]+offset]) > cb {
															if int(im[pixel[12]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[10]+offset]) > cb {
													if int(im[pixel[11]+offset]) > cb {
														if int(im[pixel[12]+offset]) > cb {
															if int(im[pixel[13]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											if int(im[pixel[10]+offset]) > cb {
												if int(im[pixel[11]+offset]) > cb {
													if int(im[pixel[12]+offset]) > cb {
														if int(im[pixel[13]+offset]) > cb {
															if int(im[pixel[14]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										}
									} else {
										if int(im[pixel[10]+offset]) > cb {
											if int(im[pixel[11]+offset]) > cb {
												if int(im[pixel[12]+offset]) > cb {
													if int(im[pixel[13]+offset]) > cb {
														if int(im[pixel[14]+offset]) > cb {
															if int(im[pixel[15]+offset]) > cb {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											continue
										}
									}
								} else {
									continue
								}
							} else {
								continue
							}
						} else {
							if int(im[pixel[8]+offset]) < c_b {
								if int(im[pixel[9]+offset]) < c_b {
									if int(im[pixel[10]+offset]) < c_b {
										if int(im[pixel[11]+offset]) < c_b {
											if int(im[pixel[12]+offset]) < c_b {
												if int(im[pixel[13]+offset]) < c_b {
													if int(im[pixel[14]+offset]) < c_b {
														if int(im[pixel[15]+offset]) < c_b {
														} else {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[4]+offset]) < c_b {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[3]+offset]) < c_b {
													if int(im[pixel[4]+offset]) < c_b {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											if int(im[pixel[2]+offset]) < c_b {
												if int(im[pixel[3]+offset]) < c_b {
													if int(im[pixel[4]+offset]) < c_b {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[6]+offset]) < c_b {
																if int(im[pixel[7]+offset]) < c_b {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										}
									} else {
										continue
									}
								} else {
									continue
								}
							} else {
								continue
							}
						}
					} else {
						if int(im[pixel[1]+offset]) < c_b {
							if int(im[pixel[2]+offset]) > cb {
								if int(im[pixel[9]+offset]) > cb {
									if int(im[pixel[7]+offset]) > cb {
										if int(im[pixel[8]+offset]) > cb {
											if int(im[pixel[10]+offset]) > cb {
												if int(im[pixel[6]+offset]) > cb {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[4]+offset]) > cb {
															if int(im[pixel[3]+offset]) > cb {
															} else {
																if int(im[pixel[11]+offset]) > cb {
																	if int(im[pixel[12]+offset]) > cb {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														} else {
															if int(im[pixel[11]+offset]) > cb {
																if int(im[pixel[12]+offset]) > cb {
																	if int(im[pixel[13]+offset]) > cb {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														if int(im[pixel[11]+offset]) > cb {
															if int(im[pixel[12]+offset]) > cb {
																if int(im[pixel[13]+offset]) > cb {
																	if int(im[pixel[14]+offset]) > cb {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[11]+offset]) > cb {
														if int(im[pixel[12]+offset]) > cb {
															if int(im[pixel[13]+offset]) > cb {
																if int(im[pixel[14]+offset]) > cb {
																	if int(im[pixel[15]+offset]) > cb {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												continue
											}
										} else {
											continue
										}
									} else {
										continue
									}
								} else {
									if int(im[pixel[9]+offset]) < c_b {
										if int(im[pixel[10]+offset]) < c_b {
											if int(im[pixel[11]+offset]) < c_b {
												if int(im[pixel[12]+offset]) < c_b {
													if int(im[pixel[13]+offset]) < c_b {
														if int(im[pixel[14]+offset]) < c_b {
															if int(im[pixel[15]+offset]) < c_b {
															} else {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																		if int(im[pixel[8]+offset]) < c_b {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														} else {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																		if int(im[pixel[8]+offset]) < c_b {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														if int(im[pixel[4]+offset]) < c_b {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																		if int(im[pixel[8]+offset]) < c_b {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[3]+offset]) < c_b {
														if int(im[pixel[4]+offset]) < c_b {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																		if int(im[pixel[8]+offset]) < c_b {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												continue
											}
										} else {
											continue
										}
									} else {
										continue
									}
								}
							} else {
								if int(im[pixel[2]+offset]) < c_b {
									if int(im[pixel[3]+offset]) > cb {
										if int(im[pixel[10]+offset]) > cb {
											if int(im[pixel[7]+offset]) > cb {
												if int(im[pixel[8]+offset]) > cb {
													if int(im[pixel[9]+offset]) > cb {
														if int(im[pixel[11]+offset]) > cb {
															if int(im[pixel[6]+offset]) > cb {
																if int(im[pixel[5]+offset]) > cb {
																	if int(im[pixel[4]+offset]) > cb {
																	} else {
																		if int(im[pixel[12]+offset]) > cb {
																			if int(im[pixel[13]+offset]) > cb {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	}
																} else {
																	if int(im[pixel[12]+offset]) > cb {
																		if int(im[pixel[13]+offset]) > cb {
																			if int(im[pixel[14]+offset]) > cb {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																}
															} else {
																if int(im[pixel[12]+offset]) > cb {
																	if int(im[pixel[13]+offset]) > cb {
																		if int(im[pixel[14]+offset]) > cb {
																			if int(im[pixel[15]+offset]) > cb {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											if int(im[pixel[10]+offset]) < c_b {
												if int(im[pixel[11]+offset]) < c_b {
													if int(im[pixel[12]+offset]) < c_b {
														if int(im[pixel[13]+offset]) < c_b {
															if int(im[pixel[14]+offset]) < c_b {
																if int(im[pixel[15]+offset]) < c_b {
																} else {
																	if int(im[pixel[6]+offset]) < c_b {
																		if int(im[pixel[7]+offset]) < c_b {
																			if int(im[pixel[8]+offset]) < c_b {
																				if int(im[pixel[9]+offset]) < c_b {
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																}
															} else {
																if int(im[pixel[5]+offset]) < c_b {
																	if int(im[pixel[6]+offset]) < c_b {
																		if int(im[pixel[7]+offset]) < c_b {
																			if int(im[pixel[8]+offset]) < c_b {
																				if int(im[pixel[9]+offset]) < c_b {
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														} else {
															if int(im[pixel[4]+offset]) < c_b {
																if int(im[pixel[5]+offset]) < c_b {
																	if int(im[pixel[6]+offset]) < c_b {
																		if int(im[pixel[7]+offset]) < c_b {
																			if int(im[pixel[8]+offset]) < c_b {
																				if int(im[pixel[9]+offset]) < c_b {
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										}
									} else {
										if int(im[pixel[3]+offset]) < c_b {
											if int(im[pixel[4]+offset]) > cb {
												if int(im[pixel[13]+offset]) > cb {
													if int(im[pixel[7]+offset]) > cb {
														if int(im[pixel[8]+offset]) > cb {
															if int(im[pixel[9]+offset]) > cb {
																if int(im[pixel[10]+offset]) > cb {
																	if int(im[pixel[11]+offset]) > cb {
																		if int(im[pixel[12]+offset]) > cb {
																			if int(im[pixel[6]+offset]) > cb {
																				if int(im[pixel[5]+offset]) > cb {
																				} else {
																					if int(im[pixel[14]+offset]) > cb {
																					} else {
																						continue
																					}
																				}
																			} else {
																				if int(im[pixel[14]+offset]) > cb {
																					if int(im[pixel[15]+offset]) > cb {
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													if int(im[pixel[13]+offset]) < c_b {
														if int(im[pixel[11]+offset]) > cb {
															if int(im[pixel[5]+offset]) > cb {
																if int(im[pixel[6]+offset]) > cb {
																	if int(im[pixel[7]+offset]) > cb {
																		if int(im[pixel[8]+offset]) > cb {
																			if int(im[pixel[9]+offset]) > cb {
																				if int(im[pixel[10]+offset]) > cb {
																					if int(im[pixel[12]+offset]) > cb {
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															if int(im[pixel[11]+offset]) < c_b {
																if int(im[pixel[12]+offset]) < c_b {
																	if int(im[pixel[14]+offset]) < c_b {
																		if int(im[pixel[15]+offset]) < c_b {
																		} else {
																			if int(im[pixel[6]+offset]) < c_b {
																				if int(im[pixel[7]+offset]) < c_b {
																					if int(im[pixel[8]+offset]) < c_b {
																						if int(im[pixel[9]+offset]) < c_b {
																							if int(im[pixel[10]+offset]) < c_b {
																							} else {
																								continue
																							}
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		}
																	} else {
																		if int(im[pixel[5]+offset]) < c_b {
																			if int(im[pixel[6]+offset]) < c_b {
																				if int(im[pixel[7]+offset]) < c_b {
																					if int(im[pixel[8]+offset]) < c_b {
																						if int(im[pixel[9]+offset]) < c_b {
																							if int(im[pixel[10]+offset]) < c_b {
																							} else {
																								continue
																							}
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														if int(im[pixel[5]+offset]) > cb {
															if int(im[pixel[6]+offset]) > cb {
																if int(im[pixel[7]+offset]) > cb {
																	if int(im[pixel[8]+offset]) > cb {
																		if int(im[pixel[9]+offset]) > cb {
																			if int(im[pixel[10]+offset]) > cb {
																				if int(im[pixel[11]+offset]) > cb {
																					if int(im[pixel[12]+offset]) > cb {
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												}
											} else {
												if int(im[pixel[4]+offset]) < c_b {
													if int(im[pixel[5]+offset]) > cb {
														if int(im[pixel[14]+offset]) > cb {
															if int(im[pixel[7]+offset]) > cb {
																if int(im[pixel[8]+offset]) > cb {
																	if int(im[pixel[9]+offset]) > cb {
																		if int(im[pixel[10]+offset]) > cb {
																			if int(im[pixel[11]+offset]) > cb {
																				if int(im[pixel[12]+offset]) > cb {
																					if int(im[pixel[13]+offset]) > cb {
																						if int(im[pixel[6]+offset]) > cb {
																						} else {
																							if int(im[pixel[15]+offset]) > cb {
																							} else {
																								continue
																							}
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															if int(im[pixel[14]+offset]) < c_b {
																if int(im[pixel[12]+offset]) > cb {
																	if int(im[pixel[6]+offset]) > cb {
																		if int(im[pixel[7]+offset]) > cb {
																			if int(im[pixel[8]+offset]) > cb {
																				if int(im[pixel[9]+offset]) > cb {
																					if int(im[pixel[10]+offset]) > cb {
																						if int(im[pixel[11]+offset]) > cb {
																							if int(im[pixel[13]+offset]) > cb {
																							} else {
																								continue
																							}
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	if int(im[pixel[12]+offset]) < c_b {
																		if int(im[pixel[13]+offset]) < c_b {
																			if int(im[pixel[15]+offset]) < c_b {
																			} else {
																				if int(im[pixel[6]+offset]) < c_b {
																					if int(im[pixel[7]+offset]) < c_b {
																						if int(im[pixel[8]+offset]) < c_b {
																							if int(im[pixel[9]+offset]) < c_b {
																								if int(im[pixel[10]+offset]) < c_b {
																									if int(im[pixel[11]+offset]) < c_b {
																									} else {
																										continue
																									}
																								} else {
																									continue
																								}
																							} else {
																								continue
																							}
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																}
															} else {
																if int(im[pixel[6]+offset]) > cb {
																	if int(im[pixel[7]+offset]) > cb {
																		if int(im[pixel[8]+offset]) > cb {
																			if int(im[pixel[9]+offset]) > cb {
																				if int(im[pixel[10]+offset]) > cb {
																					if int(im[pixel[11]+offset]) > cb {
																						if int(im[pixel[12]+offset]) > cb {
																							if int(im[pixel[13]+offset]) > cb {
																							} else {
																								continue
																							}
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														}
													} else {
														if int(im[pixel[5]+offset]) < c_b {
															if int(im[pixel[6]+offset]) > cb {
																if int(im[pixel[15]+offset]) < c_b {
																	if int(im[pixel[13]+offset]) > cb {
																		if int(im[pixel[7]+offset]) > cb {
																			if int(im[pixel[8]+offset]) > cb {
																				if int(im[pixel[9]+offset]) > cb {
																					if int(im[pixel[10]+offset]) > cb {
																						if int(im[pixel[11]+offset]) > cb {
																							if int(im[pixel[12]+offset]) > cb {
																								if int(im[pixel[14]+offset]) > cb {
																								} else {
																									continue
																								}
																							} else {
																								continue
																							}
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		if int(im[pixel[13]+offset]) < c_b {
																			if int(im[pixel[14]+offset]) < c_b {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	}
																} else {
																	if int(im[pixel[7]+offset]) > cb {
																		if int(im[pixel[8]+offset]) > cb {
																			if int(im[pixel[9]+offset]) > cb {
																				if int(im[pixel[10]+offset]) > cb {
																					if int(im[pixel[11]+offset]) > cb {
																						if int(im[pixel[12]+offset]) > cb {
																							if int(im[pixel[13]+offset]) > cb {
																								if int(im[pixel[14]+offset]) > cb {
																								} else {
																									continue
																								}
																							} else {
																								continue
																							}
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																}
															} else {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) > cb {
																		if int(im[pixel[14]+offset]) > cb {
																			if int(im[pixel[8]+offset]) > cb {
																				if int(im[pixel[9]+offset]) > cb {
																					if int(im[pixel[10]+offset]) > cb {
																						if int(im[pixel[11]+offset]) > cb {
																							if int(im[pixel[12]+offset]) > cb {
																								if int(im[pixel[13]+offset]) > cb {
																									if int(im[pixel[15]+offset]) > cb {
																									} else {
																										continue
																									}
																								} else {
																									continue
																								}
																							} else {
																								continue
																							}
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			if int(im[pixel[14]+offset]) < c_b {
																				if int(im[pixel[15]+offset]) < c_b {
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		}
																	} else {
																		if int(im[pixel[7]+offset]) < c_b {
																			if int(im[pixel[8]+offset]) < c_b {
																			} else {
																				if int(im[pixel[15]+offset]) < c_b {
																				} else {
																					continue
																				}
																			}
																		} else {
																			if int(im[pixel[14]+offset]) < c_b {
																				if int(im[pixel[15]+offset]) < c_b {
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		}
																	}
																} else {
																	if int(im[pixel[13]+offset]) > cb {
																		if int(im[pixel[7]+offset]) > cb {
																			if int(im[pixel[8]+offset]) > cb {
																				if int(im[pixel[9]+offset]) > cb {
																					if int(im[pixel[10]+offset]) > cb {
																						if int(im[pixel[11]+offset]) > cb {
																							if int(im[pixel[12]+offset]) > cb {
																								if int(im[pixel[14]+offset]) > cb {
																									if int(im[pixel[15]+offset]) > cb {
																									} else {
																										continue
																									}
																								} else {
																									continue
																								}
																							} else {
																								continue
																							}
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		if int(im[pixel[13]+offset]) < c_b {
																			if int(im[pixel[14]+offset]) < c_b {
																				if int(im[pixel[15]+offset]) < c_b {
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	}
																}
															}
														} else {
															if int(im[pixel[12]+offset]) > cb {
																if int(im[pixel[7]+offset]) > cb {
																	if int(im[pixel[8]+offset]) > cb {
																		if int(im[pixel[9]+offset]) > cb {
																			if int(im[pixel[10]+offset]) > cb {
																				if int(im[pixel[11]+offset]) > cb {
																					if int(im[pixel[13]+offset]) > cb {
																						if int(im[pixel[14]+offset]) > cb {
																							if int(im[pixel[6]+offset]) > cb {
																							} else {
																								if int(im[pixel[15]+offset]) > cb {
																								} else {
																									continue
																								}
																							}
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																if int(im[pixel[12]+offset]) < c_b {
																	if int(im[pixel[13]+offset]) < c_b {
																		if int(im[pixel[14]+offset]) < c_b {
																			if int(im[pixel[15]+offset]) < c_b {
																			} else {
																				if int(im[pixel[6]+offset]) < c_b {
																					if int(im[pixel[7]+offset]) < c_b {
																						if int(im[pixel[8]+offset]) < c_b {
																							if int(im[pixel[9]+offset]) < c_b {
																								if int(im[pixel[10]+offset]) < c_b {
																									if int(im[pixel[11]+offset]) < c_b {
																									} else {
																										continue
																									}
																								} else {
																									continue
																								}
																							} else {
																								continue
																							}
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														}
													}
												} else {
													if int(im[pixel[11]+offset]) > cb {
														if int(im[pixel[7]+offset]) > cb {
															if int(im[pixel[8]+offset]) > cb {
																if int(im[pixel[9]+offset]) > cb {
																	if int(im[pixel[10]+offset]) > cb {
																		if int(im[pixel[12]+offset]) > cb {
																			if int(im[pixel[13]+offset]) > cb {
																				if int(im[pixel[6]+offset]) > cb {
																					if int(im[pixel[5]+offset]) > cb {
																					} else {
																						if int(im[pixel[14]+offset]) > cb {
																						} else {
																							continue
																						}
																					}
																				} else {
																					if int(im[pixel[14]+offset]) > cb {
																						if int(im[pixel[15]+offset]) > cb {
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														if int(im[pixel[11]+offset]) < c_b {
															if int(im[pixel[12]+offset]) < c_b {
																if int(im[pixel[13]+offset]) < c_b {
																	if int(im[pixel[14]+offset]) < c_b {
																		if int(im[pixel[15]+offset]) < c_b {
																		} else {
																			if int(im[pixel[6]+offset]) < c_b {
																				if int(im[pixel[7]+offset]) < c_b {
																					if int(im[pixel[8]+offset]) < c_b {
																						if int(im[pixel[9]+offset]) < c_b {
																							if int(im[pixel[10]+offset]) < c_b {
																							} else {
																								continue
																							}
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		}
																	} else {
																		if int(im[pixel[5]+offset]) < c_b {
																			if int(im[pixel[6]+offset]) < c_b {
																				if int(im[pixel[7]+offset]) < c_b {
																					if int(im[pixel[8]+offset]) < c_b {
																						if int(im[pixel[9]+offset]) < c_b {
																							if int(im[pixel[10]+offset]) < c_b {
																							} else {
																								continue
																							}
																						} else {
																							continue
																						}
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												}
											}
										} else {
											if int(im[pixel[10]+offset]) > cb {
												if int(im[pixel[7]+offset]) > cb {
													if int(im[pixel[8]+offset]) > cb {
														if int(im[pixel[9]+offset]) > cb {
															if int(im[pixel[11]+offset]) > cb {
																if int(im[pixel[12]+offset]) > cb {
																	if int(im[pixel[6]+offset]) > cb {
																		if int(im[pixel[5]+offset]) > cb {
																			if int(im[pixel[4]+offset]) > cb {
																			} else {
																				if int(im[pixel[13]+offset]) > cb {
																				} else {
																					continue
																				}
																			}
																		} else {
																			if int(im[pixel[13]+offset]) > cb {
																				if int(im[pixel[14]+offset]) > cb {
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		}
																	} else {
																		if int(im[pixel[13]+offset]) > cb {
																			if int(im[pixel[14]+offset]) > cb {
																				if int(im[pixel[15]+offset]) > cb {
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												if int(im[pixel[10]+offset]) < c_b {
													if int(im[pixel[11]+offset]) < c_b {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[13]+offset]) < c_b {
																if int(im[pixel[14]+offset]) < c_b {
																	if int(im[pixel[15]+offset]) < c_b {
																	} else {
																		if int(im[pixel[6]+offset]) < c_b {
																			if int(im[pixel[7]+offset]) < c_b {
																				if int(im[pixel[8]+offset]) < c_b {
																					if int(im[pixel[9]+offset]) < c_b {
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	}
																} else {
																	if int(im[pixel[5]+offset]) < c_b {
																		if int(im[pixel[6]+offset]) < c_b {
																			if int(im[pixel[7]+offset]) < c_b {
																				if int(im[pixel[8]+offset]) < c_b {
																					if int(im[pixel[9]+offset]) < c_b {
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																}
															} else {
																if int(im[pixel[4]+offset]) < c_b {
																	if int(im[pixel[5]+offset]) < c_b {
																		if int(im[pixel[6]+offset]) < c_b {
																			if int(im[pixel[7]+offset]) < c_b {
																				if int(im[pixel[8]+offset]) < c_b {
																					if int(im[pixel[9]+offset]) < c_b {
																					} else {
																						continue
																					}
																				} else {
																					continue
																				}
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										}
									}
								} else {
									if int(im[pixel[9]+offset]) > cb {
										if int(im[pixel[7]+offset]) > cb {
											if int(im[pixel[8]+offset]) > cb {
												if int(im[pixel[10]+offset]) > cb {
													if int(im[pixel[11]+offset]) > cb {
														if int(im[pixel[6]+offset]) > cb {
															if int(im[pixel[5]+offset]) > cb {
																if int(im[pixel[4]+offset]) > cb {
																	if int(im[pixel[3]+offset]) > cb {
																	} else {
																		if int(im[pixel[12]+offset]) > cb {
																		} else {
																			continue
																		}
																	}
																} else {
																	if int(im[pixel[12]+offset]) > cb {
																		if int(im[pixel[13]+offset]) > cb {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																}
															} else {
																if int(im[pixel[12]+offset]) > cb {
																	if int(im[pixel[13]+offset]) > cb {
																		if int(im[pixel[14]+offset]) > cb {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														} else {
															if int(im[pixel[12]+offset]) > cb {
																if int(im[pixel[13]+offset]) > cb {
																	if int(im[pixel[14]+offset]) > cb {
																		if int(im[pixel[15]+offset]) > cb {
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											continue
										}
									} else {
										if int(im[pixel[9]+offset]) < c_b {
											if int(im[pixel[10]+offset]) < c_b {
												if int(im[pixel[11]+offset]) < c_b {
													if int(im[pixel[12]+offset]) < c_b {
														if int(im[pixel[13]+offset]) < c_b {
															if int(im[pixel[14]+offset]) < c_b {
																if int(im[pixel[15]+offset]) < c_b {
																} else {
																	if int(im[pixel[6]+offset]) < c_b {
																		if int(im[pixel[7]+offset]) < c_b {
																			if int(im[pixel[8]+offset]) < c_b {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																}
															} else {
																if int(im[pixel[5]+offset]) < c_b {
																	if int(im[pixel[6]+offset]) < c_b {
																		if int(im[pixel[7]+offset]) < c_b {
																			if int(im[pixel[8]+offset]) < c_b {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														} else {
															if int(im[pixel[4]+offset]) < c_b {
																if int(im[pixel[5]+offset]) < c_b {
																	if int(im[pixel[6]+offset]) < c_b {
																		if int(im[pixel[7]+offset]) < c_b {
																			if int(im[pixel[8]+offset]) < c_b {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														if int(im[pixel[3]+offset]) < c_b {
															if int(im[pixel[4]+offset]) < c_b {
																if int(im[pixel[5]+offset]) < c_b {
																	if int(im[pixel[6]+offset]) < c_b {
																		if int(im[pixel[7]+offset]) < c_b {
																			if int(im[pixel[8]+offset]) < c_b {
																			} else {
																				continue
																			}
																		} else {
																			continue
																		}
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											continue
										}
									}
								}
							}
						} else {
							if int(im[pixel[8]+offset]) > cb {
								if int(im[pixel[7]+offset]) > cb {
									if int(im[pixel[9]+offset]) > cb {
										if int(im[pixel[10]+offset]) > cb {
											if int(im[pixel[6]+offset]) > cb {
												if int(im[pixel[5]+offset]) > cb {
													if int(im[pixel[4]+offset]) > cb {
														if int(im[pixel[3]+offset]) > cb {
															if int(im[pixel[2]+offset]) > cb {
															} else {
																if int(im[pixel[11]+offset]) > cb {
																} else {
																	continue
																}
															}
														} else {
															if int(im[pixel[11]+offset]) > cb {
																if int(im[pixel[12]+offset]) > cb {
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														if int(im[pixel[11]+offset]) > cb {
															if int(im[pixel[12]+offset]) > cb {
																if int(im[pixel[13]+offset]) > cb {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[11]+offset]) > cb {
														if int(im[pixel[12]+offset]) > cb {
															if int(im[pixel[13]+offset]) > cb {
																if int(im[pixel[14]+offset]) > cb {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[11]+offset]) > cb {
													if int(im[pixel[12]+offset]) > cb {
														if int(im[pixel[13]+offset]) > cb {
															if int(im[pixel[14]+offset]) > cb {
																if int(im[pixel[15]+offset]) > cb {
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											continue
										}
									} else {
										continue
									}
								} else {
									continue
								}
							} else {
								if int(im[pixel[8]+offset]) < c_b {
									if int(im[pixel[9]+offset]) < c_b {
										if int(im[pixel[10]+offset]) < c_b {
											if int(im[pixel[11]+offset]) < c_b {
												if int(im[pixel[12]+offset]) < c_b {
													if int(im[pixel[13]+offset]) < c_b {
														if int(im[pixel[14]+offset]) < c_b {
															if int(im[pixel[15]+offset]) < c_b {
															} else {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															}
														} else {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														}
													} else {
														if int(im[pixel[4]+offset]) < c_b {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[3]+offset]) < c_b {
														if int(im[pixel[4]+offset]) < c_b {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[2]+offset]) < c_b {
													if int(im[pixel[3]+offset]) < c_b {
														if int(im[pixel[4]+offset]) < c_b {
															if int(im[pixel[5]+offset]) < c_b {
																if int(im[pixel[6]+offset]) < c_b {
																	if int(im[pixel[7]+offset]) < c_b {
																	} else {
																		continue
																	}
																} else {
																	continue
																}
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											continue
										}
									} else {
										continue
									}
								} else {
									continue
								}
							}
						}
					}
				} else {
					if int(im[pixel[7]+offset]) > cb {
						if int(im[pixel[8]+offset]) > cb {
							if int(im[pixel[9]+offset]) > cb {
								if int(im[pixel[6]+offset]) > cb {
									if int(im[pixel[5]+offset]) > cb {
										if int(im[pixel[4]+offset]) > cb {
											if int(im[pixel[3]+offset]) > cb {
												if int(im[pixel[2]+offset]) > cb {
													if int(im[pixel[1]+offset]) > cb {
													} else {
														if int(im[pixel[10]+offset]) > cb {
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[10]+offset]) > cb {
														if int(im[pixel[11]+offset]) > cb {
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[10]+offset]) > cb {
													if int(im[pixel[11]+offset]) > cb {
														if int(im[pixel[12]+offset]) > cb {
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											if int(im[pixel[10]+offset]) > cb {
												if int(im[pixel[11]+offset]) > cb {
													if int(im[pixel[12]+offset]) > cb {
														if int(im[pixel[13]+offset]) > cb {
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										}
									} else {
										if int(im[pixel[10]+offset]) > cb {
											if int(im[pixel[11]+offset]) > cb {
												if int(im[pixel[12]+offset]) > cb {
													if int(im[pixel[13]+offset]) > cb {
														if int(im[pixel[14]+offset]) > cb {
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											continue
										}
									}
								} else {
									if int(im[pixel[10]+offset]) > cb {
										if int(im[pixel[11]+offset]) > cb {
											if int(im[pixel[12]+offset]) > cb {
												if int(im[pixel[13]+offset]) > cb {
													if int(im[pixel[14]+offset]) > cb {
														if int(im[pixel[15]+offset]) > cb {
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											continue
										}
									} else {
										continue
									}
								}
							} else {
								continue
							}
						} else {
							continue
						}
					} else {
						if int(im[pixel[7]+offset]) < c_b {
							if int(im[pixel[8]+offset]) < c_b {
								if int(im[pixel[9]+offset]) < c_b {
									if int(im[pixel[6]+offset]) < c_b {
										if int(im[pixel[5]+offset]) < c_b {
											if int(im[pixel[4]+offset]) < c_b {
												if int(im[pixel[3]+offset]) < c_b {
													if int(im[pixel[2]+offset]) < c_b {
														if int(im[pixel[1]+offset]) < c_b {
														} else {
															if int(im[pixel[10]+offset]) < c_b {
															} else {
																continue
															}
														}
													} else {
														if int(im[pixel[10]+offset]) < c_b {
															if int(im[pixel[11]+offset]) < c_b {
															} else {
																continue
															}
														} else {
															continue
														}
													}
												} else {
													if int(im[pixel[10]+offset]) < c_b {
														if int(im[pixel[11]+offset]) < c_b {
															if int(im[pixel[12]+offset]) < c_b {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												}
											} else {
												if int(im[pixel[10]+offset]) < c_b {
													if int(im[pixel[11]+offset]) < c_b {
														if int(im[pixel[12]+offset]) < c_b {
															if int(im[pixel[13]+offset]) < c_b {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											}
										} else {
											if int(im[pixel[10]+offset]) < c_b {
												if int(im[pixel[11]+offset]) < c_b {
													if int(im[pixel[12]+offset]) < c_b {
														if int(im[pixel[13]+offset]) < c_b {
															if int(im[pixel[14]+offset]) < c_b {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										}
									} else {
										if int(im[pixel[10]+offset]) < c_b {
											if int(im[pixel[11]+offset]) < c_b {
												if int(im[pixel[12]+offset]) < c_b {
													if int(im[pixel[13]+offset]) < c_b {
														if int(im[pixel[14]+offset]) < c_b {
															if int(im[pixel[15]+offset]) < c_b {
															} else {
																continue
															}
														} else {
															continue
														}
													} else {
														continue
													}
												} else {
													continue
												}
											} else {
												continue
											}
										} else {
											continue
										}
									}
								} else {
									continue
								}
							} else {
								continue
							}
						} else {
							continue
						}
					}
				}
			}
			ret_corners = append(ret_corners, image.Point{X: x, Y: y})
		}
	}
	return ret_corners
}
